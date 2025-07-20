package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/h-varmazyar/Gate/services/gather/api/rest/v1"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type Router struct {
	log      *log.Logger
	v1Router v1.Router
}

func New(log *log.Logger, v1Router v1.Router) Router {
	return Router{
		v1Router: v1Router,
		log:      log,
	}
}

func (r Router) StartServing(ginEngine *gin.Engine, address string) error {
	r.RegisterRoutes(ginEngine)

	srv := &http.Server{
		Addr:    address,
		Handler: ginEngine,
	}

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		fmt.Println("[My Demo] Failed to start HTTP Server at", srv.Addr)
		return err
	}
	go srv.Serve(ln)

	return err
}

func (r Router) RegisterRoutes(ginRouter *gin.Engine) {
	r.log.Infof("********** registering routes")
	apiRouter := ginRouter.Group("/api")
	r.v1Router.RegisterRoutes(apiRouter)
}
