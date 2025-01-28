package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/Gate/services/gather/api/rest/v1/candles"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Router struct {
	logger         *log.Logger
	candlesHandler candles.Handler
}

func NewRouter(logger *log.Logger, candlesHandler candles.Handler) *Router {
	return &Router{
		logger:         logger,
		candlesHandler: candlesHandler,
	}
}

func (r *Router) RegisterRoutes(ginRouter *gin.RouterGroup) {
	r.logger.Infof("********** registering v1 routes")
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
	apiRouter := ginRouter.Group("/v1")

	candlesRouter := apiRouter.Group("/candles")
	candlesRouter.GET("/markets/:market-id/resolutions/:resolution-id", r.candlesHandler.List)
}
