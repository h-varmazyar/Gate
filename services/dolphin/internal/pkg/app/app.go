package app

import (
	"fmt"
	"github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/defaults"
	"github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/httpx"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/sessions"

	"github.com/gorilla/mux"
)

// App is where it all happens! It holds on to options,
// the underlying router, the middleware, and more.
// Without an App you can't do much!
type App struct {
	Options
	// Middleware returns the current MiddlewareStack for the App/Group.
	Middleware    *MiddlewareStack `json:"-"`
	ErrorHandlers ErrorHandlers    `json:"-"`
	router        *mux.Router
	moot          *sync.RWMutex
	routes        RouteList
	root          *App
	store         *sessions.CookieStore
	children      []*App
	filepaths     []string
}

// Muxer returns the underlying mux router to allow
// for advance configurations
func (a *App) Muxer() *mux.Router {
	return a.router
}

// New returns a new instance of App and adds some sane, and useful, defaults.
func New(opts Options) *App {

	opts = optionsWithDefaults(opts)

	a := &App{
		Options:  opts,
		router:   mux.NewRouter(),
		moot:     &sync.RWMutex{},
		store:    sessions.NewCookieStore([]byte("h081dpYowvPrDNGa4REV84AodtSoJdUBUxXGzG7QnZ8vHDdNONDHBqjF1bGEEyvS")),
		routes:   RouteList{},
		children: []*App{},
	}
	dem := a.defaultErrorMiddleware
	a.Middleware = newMiddlewareStack(dem)

	notFoundHandler := func(errorf string, code int) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			c := a.newContext(RouteInfo{}, res, req)
			err := fmt.Errorf(errorf, req.Method, req.URL.Path)
			_ = a.ErrorHandlers.Get(code)(code, err, c)
		}
	}

	a.router.NotFoundHandler = notFoundHandler("path not found: %s %s", http.StatusNotFound)
	a.router.MethodNotAllowedHandler = notFoundHandler("method not found: %s %s", http.StatusMethodNotAllowed)

	if a.MethodOverride == nil {
		a.MethodOverride = func(res http.ResponseWriter, req *http.Request) {
			if req.Method == "POST" {
				req.Method = defaults.String(req.FormValue("_method"), "POST")
				req.Form.Del("_method")
				req.PostForm.Del("_method")
			}
		}
	}

	a.Use(a.PanicHandler)
	//a.Use(RequestLogger)
	//a.Use(sessionSaver)

	return a
}

func (a *App) newContext(info RouteInfo, res http.ResponseWriter, req *http.Request) Context {
	if ws, ok := res.(*Response); ok {
		res = ws
	}

	// Parse URL Params
	params := url.Values{}
	vars := mux.Vars(req)
	for k, v := range vars {
		params.Add(k, v)
	}

	// Parse URL Query String Params
	// For POST, PUT, and PATCH requests, it also parse the request body as a form.
	// Request body parameters take precedence over URL query string values in params
	if err := req.ParseForm(); err == nil {
		for k, v := range req.Form {
			for _, vv := range v {
				params.Add(k, vv)
			}
		}
	}

	ct := httpx.ContentType(req)

	data := &sync.Map{}

	data.Store("app", a)
	data.Store("env", a.Env)
	data.Store("routes", a.Routes())
	data.Store("current_route", info)
	data.Store("current_path", req.URL.Path)
	data.Store("contentType", ct)
	data.Store("method", req.Method)

	for _, route := range a.Routes() {
		cRoute := route
		data.Store(cRoute.PathName, cRoute.BuildPathHelper())
	}

	return &AppContext{
		Context:      req.Context(),
		response:     res,
		request:      req,
		params:       params,
		logger:       a.Logger,
		accessLogger: a.AccessLogger,
		store:        a.store,
		contentType:  ct,
		data:         data,
	}
}

//func RequestLogger(h Handler) Handler {
//	return func(c Context) error {
//		rs, err := randString(10)
//		if err != nil {
//			return err
//		}
//		var irid interface{}
//		if irid = c.Session().Get("requestor_id"); irid == nil {
//			rs, err := randString(10)
//			if err != nil {
//				return err
//			}
//			irid = rs
//			c.Session().Set("requestor_id", irid)
//			c.Session().Save()
//		}
//
//		rid := irid.(string) + "-" + rs
//		c.Set("request_id", rid)
//		c.LogField("request_id", rid)
//
//		start := time.Now()
//		defer func() {
//			ws, ok := c.Response().(*Response)
//			if !ok {
//				ws = &Response{ResponseWriter: c.Response()}
//				ws.Status = http.StatusOK
//			}
//			req := c.Request()
//			ct := httpx.ContentType(req)
//			if ct != "" {
//				c.LogField("content_type", ct)
//			}
//			c.LogFields(map[string]interface{}{
//				"method":     req.Method,
//				"path":       req.URL.String(),
//				"duration":   time.Since(start),
//				"size":       ws.Size,
//				"human_size": humanize.Bytes(uint64(ws.Size)),
//				"status":     ws.Status,
//			})
//			c.Logger().Info(req.URL.String())
//		}()
//		return h(c)
//	}
//}
