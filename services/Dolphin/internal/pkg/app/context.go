package app

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Context holds on to information as you
// pass it down through middleware, Handlers,
// templates, etc... It strives to make your
// life a happier one.
type Context interface {
	context.Context
	Response() http.ResponseWriter
	Request() *http.Request
	Session() *sessions.Session
	//Cookies() *Cookies
	Params() ParamValues
	Param(string) string
	Set(string, interface{})
	LogField(string, interface{})
	LogFields(map[string]interface{})
	Logger() *logrus.Logger
	AccessLogger() *logrus.Logger
	NewAccessLog(operation string, extraInfo interface{})
	Bind(interface{}) error
	Render(int, string, ...TemplateFunc) error
	Error(int, error) error
	Redirect(string) error
	RedirectStatus(int, string) error
	Data() map[string]interface{}
	//Flash() *Flash
	//File(string) (binding.File, error)
}

// ParamValues will most commonly be url.Values,
// but isn't it great that you set your own? :)
type ParamValues interface {
	Get(string) string
}

type TemplateFunc struct {
	Key  string
	Func interface{}
}
