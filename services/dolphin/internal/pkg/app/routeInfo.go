package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 14.11.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

type RouteInfo struct {
	Method       string     `json:"method"`
	Path         string     `json:"path"`
	HandlerName  string     `json:"handler"`
	ResourceName string     `json:"resourceName,omitempty"`
	PathName     string     `json:"pathName"`
	Aliases      []string   `json:"aliases"`
	MuxRoute     *mux.Route `json:"-"`
	Handler      Handler    `json:"-"`
	App          *App       `json:"-"`
}

// String returns a JSON representation of the RouteInfo
func (ri RouteInfo) String() string {
	b, _ := json.MarshalIndent(ri, "", "  ")
	return string(b)
}

// Alias path patterns to the this route. This is not the
// same as a redirect.
func (ri *RouteInfo) Alias(aliases ...string) *RouteInfo {
	ri.Aliases = append(ri.Aliases, aliases...)
	for _, a := range aliases {
		ri.App.router.Handle(a, ri).Methods(ri.Method)
	}
	return ri
}

// Name allows users to set custom names for the routes.
func (ri *RouteInfo) Name(name string) *RouteInfo {
	routeIndex := -1
	for index, route := range ri.App.Routes() {
		if route.Path == ri.Path && route.Method == ri.Method {
			routeIndex = index
			break
		}
	}

	//name = flect.Camelize(name)

	if !strings.HasSuffix(name, "Path") {
		name = name + "Path"
	}

	ri.PathName = name
	if routeIndex != -1 {
		ri.App.Routes()[routeIndex] = reflect.ValueOf(ri).Interface().(*RouteInfo)
	}

	return ri
}

// BuildPathHelper Builds a routeHelperfunc for a particular RouteInfo
func (ri *RouteInfo) BuildPathHelper() RouteHelperFunc {
	cRoute := ri
	return func(opts map[string]interface{}) (template.HTML, error) {
		pairs := []string{}
		for k, v := range opts {
			pairs = append(pairs, k)
			pairs = append(pairs, fmt.Sprintf("%v", v))
		}

		url, err := cRoute.MuxRoute.URL(pairs...)
		if err != nil {
			return "", fmt.Errorf("missing parameters for %v: %s", cRoute.Path, err)
		}

		result := url.Path
		result = addExtraParamsTo(result, opts)

		return template.HTML(result), nil
	}
}

func (ri RouteInfo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	a := ri.App

	c := a.newContext(ri, res, req)
	//defer c.Flash().persist(c.Session())

	//payload := events.Payload{
	//	"route":   ri,
	//	"app":     a,
	//	"context": c,
	//}
	//
	//events.EmitPayload(EvtRouteStarted, payload)

	err := a.Middleware.handler(ri)(c)

	if err != nil {
		status := http.StatusInternalServerError
		if he, ok := err.(HTTPError); ok {
			status = he.Status
		}
		//events.EmitError(EvtRouteErr, err, payload)
		// things have really hit the fan if we're here!!
		a.Logger.Error(err)
		c.Response().WriteHeader(status)
		c.Response().Write([]byte(err.Error()))
	}
	//events.EmitPayload(EvtRouteFinished, payload)
}