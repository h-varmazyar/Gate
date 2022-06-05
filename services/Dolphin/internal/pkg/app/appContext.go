package app

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	"github.com/gorilla/sessions"
)

// assert that Context is implementing Context
var _ Context = &AppContext{}
var _ context.Context = &AppContext{}

// Context is, as its name implies, a default
// implementation of the Context interface.
type AppContext struct {
	context.Context
	response     http.ResponseWriter
	request      *http.Request
	params       url.Values
	logger       *logrus.Logger
	accessLogger *logrus.Logger
	store        *sessions.CookieStore
	contentType  string
	data         *sync.Map
	//flash       *Flash
}

// type Session struct {
// 	Session *sessions.Session
// 	req     *http.Request
// 	res     http.ResponseWriter
// }

func (d *AppContext) SetValue(key, value string) {
	d.Context = metadata.AppendToOutgoingContext(d.Context, key, value)
	d.Context = context.WithValue(d.Context, key, value)
}

func (d *AppContext) GetValue(key string) (string, bool) {
	value, ok := d.Value(key).(string)
	if ok {
		return value, true
	}
	var md metadata.MD
	md, ok = metadata.FromIncomingContext(d)
	if !ok {
		return "", false
	}
	if len(md[key]) > 0 {
		return md[key][0], true
	}
	return "", false
}

func (d *AppContext) SetToken(token string) {
	if token == "" {
		return
	}
	d.SetValue("token", token)
}

func (d *AppContext) GetToken() (string, bool) {
	value, ok := d.GetValue("token")
	if !ok {
		return "", false
	}
	return value, true
}

// Response returns the original Response for the request.
func (d *AppContext) Response() http.ResponseWriter {
	return d.response
}

// Request returns the original Request.
func (d *AppContext) Request() *http.Request {
	return d.request
}

// Params returns all of the parameters for the request,
// including both named params and query string parameters.
func (d *AppContext) Params() ParamValues {
	return d.params
}

// Logger returns the Logger for this context.
func (d *AppContext) Logger() *logrus.Logger {
	return d.logger
}

func (d *AppContext) AccessLogger() *logrus.Logger {
	return d.accessLogger
}

func (d *AppContext) NewAccessLog(operation string, extraInfo interface{}) {
	d.accessLogger.Infof("%s by %s from %s (%s) - more info: %v",
		operation, d.Session().Values["username"], d.request.RemoteAddr, d.request.Header.Get("User-Agent"), extraInfo)
}

// Param returns a param, either named or query string,
// based on the key.
func (d *AppContext) Param(key string) string {
	return d.Params().Get(key)
}

// Set a value onto the Context. Any value set onto the Context
// will be automatically available in templates.
func (d *AppContext) Set(key string, value interface{}) {
	d.data.Store(key, value)
}

// Value that has previously stored on the context.
func (d *AppContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v, ok := d.data.Load(k); ok {
			return v
		}
	}
	return d.Context.Value(key)
}

// Session for the associated Request.
func (d *AppContext) Session() *sessions.Session {
	s, _ := d.store.Get(d.Request(), "session")
	return s
}

//// Cookies for the associated request and response.
//func (d *Context) Cookies() *Cookies {
//	return &Cookies{d.request, d.response}
//}
//
//// Flash messages for the associated Request.
//func (d *Context) Flash() *Flash {
//	return d.flash
//}

//type paginable interface {
//	Paginate() string
//}

func (d *AppContext) Render(status int, name string, functions ...TemplateFunc) error {
	d.Response().WriteHeader(status)

	d.Set("URL", d.Request().URL.String())

	s := d.Session()
	d.Set("Session", s.Values)

	d.Set("Template", name)

	functionMap := template.FuncMap{}

	for _, function := range functions {
		functionMap[function.Key] = function.Func
	}
	t := template.Must(template.New("base.html").Funcs(functionMap).ParseFiles(
		"./public/views/layout/base.html",
		"./public/views/layout/header.html",
		"./public/views/layout/aside.html",
		"./public/views/"+name+".html",
	))

	return t.Execute(d.Response(), d.Data())
}

// Bind the interface to the request.Body. The type of binding
// is dependent on the "Content-Type" for the request. If the type
// is "application/json" it will use "json.NewDecoder". If the type
// is "application/xml" it will use "xml.NewDecoder". See the
// github.com/gobuffalo/buffalo/binding package for more details.
func (d *AppContext) Bind(_ interface{}) error {
	//return binding.Exec(d.Request(), value)
	return nil
}

// LogField adds the key/value pair onto the Logger to be printed out
// as part of the request logging. This allows you to easily add things
// like metrics (think DB times) to your request.
func (d *AppContext) LogField(_ string, _ interface{}) {
	//d.logger = d.logger.WithField(key, value)
}

// LogFields adds the key/value pairs onto the Logger to be printed out
// as part of the request logging. This allows you to easily add things
// like metrics (think DB times) to your request.
func (d *AppContext) LogFields(_ map[string]interface{}) {
	//d.logger = d.logger.WithFields(values)
}

func (d *AppContext) Error(status int, err error) error {
	d.Response().WriteHeader(status)
	_, _ = d.Response().Write([]byte(err.Error()))
	return nil
}

//var mapType = reflect.ValueOf(map[string]interface{}{}).Type()

func (d *AppContext) RedirectStatus(status int, url string) error {
	http.Redirect(d.Response(), d.Request(), url, status)
	return nil
}
func (d *AppContext) Redirect(url string) error {
	http.Redirect(d.Response(), d.Request(), url, 302)
	return nil
}

// Data contains all the values set through Get/Set.
func (d *AppContext) Data() map[string]interface{} {
	m := map[string]interface{}{}
	d.data.Range(func(k, v interface{}) bool {
		s, ok := k.(string)
		if !ok {
			return false
		}
		m[s] = v
		return true
	})
	return m
}

func (d *AppContext) String() string {
	data := d.Data()
	bb := make([]string, 0, len(data))

	for k, v := range data {
		if _, ok := v.(RouteHelperFunc); !ok {
			bb = append(bb, fmt.Sprintf("%s: %s", k, v))
		}
	}
	sort.Strings(bb)
	return strings.Join(bb, "\n\n")
}

//// File returns an uploaded file by name, or an error
//func (d *Context) File(name string) (binding.File, error) {
//	req := d.Request()
//	if err := req.ParseMultipartForm(5 * 1024 * 1024); err != nil {
//		return binding.File{}, err
//	}
//	f, h, err := req.FormFile(name)
//	app := binding.File{
//		File:       f,
//		FileHeader: h,
//	}
//	if err != nil {
//		return app, err
//	}
//	return app, nil
//}

// MarshalJSON implements json marshaling for the context
func (d *AppContext) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	data := d.Data()
	for k, v := range data {
		// don't try and marshal ourself
		if _, ok := v.(*Context); ok {
			continue
		}
		if _, err := json.Marshal(v); err == nil {
			// it can be marshaled, so add it:
			m[k] = v
		}
	}
	return json.Marshal(m)
}
