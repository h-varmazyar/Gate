package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"go.uber.org/fx"
	"net"
	"net/http"
)

type Router struct {
	v1 V1
}

type Params struct {
	fx.In

	GinEngine *gin.Engine
	V1        V1
}

type Result struct {
	fx.Out

	Router *Router
}

func New(lc fx.Lifecycle, params Params) Result {
	router := &Router{
		v1: params.V1,
	}
	router.RegisterRoutes(params.GinEngine)

	srv := &http.Server{
		Addr:    ":8765",
		Handler: params.GinEngine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Println("[My Demo] Failed to start HTTP Server at", srv.Addr)
				return err
			}
			go srv.Serve(ln)
			fmt.Println("[My Demo]Succeeded to start HTTP Server at", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return Result{Router: router}
}

func (r *Router) RegisterRoutes(ginRouter *gin.Engine) {
	log.Infof("********** registering routes")
	ginRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
	apiRouter := ginRouter.Group("/api")
	r.v1.RegisterRoutes(apiRouter)
}
