package core

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/core/brokerages"
	functions "github.com/h-varmazyar/Gate/services/raven/internal/app/core/functions"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/core/platforms"
	log "github.com/sirupsen/logrus"
)

func RegisterRoutes(router *gorilla.Router, logger *log.Logger, configs *Configs) {
	coreRouter := router.PathPrefix("/core").Subrouter()

	brokerages.HandlerInstance(logger, configs.CoreAddress).RegisterRoutes(coreRouter)
	functions.HandlerInstance(logger, configs.CoreAddress).RegisterRoutes(coreRouter)
	platforms.HandlerInstance(logger, configs.CoreAddress).RegisterRoutes(coreRouter)
}
