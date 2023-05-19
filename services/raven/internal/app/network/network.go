package network

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/network/ips"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/network/rateLimiters"
	log "github.com/sirupsen/logrus"
)

func RegisterRoutes(router *gorilla.Router, logger *log.Logger, configs *Configs) {
	networkRouter := router.PathPrefix("/network").Subrouter()
	ips.ControllerInstance(logger, configs.NetworkAddress).RegisterRoutes(networkRouter)
	rateLimiters.ControllerInstance(logger, configs.NetworkAddress).RegisterRoutes(networkRouter)
}
