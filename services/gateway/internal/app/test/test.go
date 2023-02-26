package test

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Controller struct {
	log *log.Logger
}

func RegisterRoutes(router *gorilla.Router, logger *log.Logger) {
	c := &Controller{log: logger}
	testRouter := router.PathPrefix("/test").Subrouter()
	testRouter.HandleFunc("/ping", c.ping).Methods(http.MethodGet)
}

func (c *Controller) ping(res http.ResponseWriter, req *http.Request) {
	jsonM := struct {
		Pong string
	}{
		Pong: time.Now().String(),
	}
	httpext.SendModel(res, req, http.StatusOK, jsonM)
}
