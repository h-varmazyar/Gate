package chipmunk

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk/assets"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk/candles"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk/indicators"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk/markets"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk/resolutions"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk/wallets"
	log "github.com/sirupsen/logrus"
)

func RegisterRoutes(router *gorilla.Router, logger *log.Logger, configs *Configs) {
	chipmunkRouter := router.PathPrefix("/chipmunk").Subrouter()
	assets.ControllerInstance(logger, configs.ChipmunkAddress).RegisterRoutes(chipmunkRouter)
	candles.ControllerInstance(logger, configs.ChipmunkAddress).RegisterRoutes(chipmunkRouter)
	indicators.ControllerInstance(logger, configs.ChipmunkAddress).RegisterRoutes(chipmunkRouter)
	markets.ControllerInstance(logger, configs.ChipmunkAddress).RegisterRoutes(chipmunkRouter)
	resolutions.ControllerInstance(logger, configs.ChipmunkAddress).RegisterRoutes(chipmunkRouter)
	wallets.ControllerInstance(logger, configs.ChipmunkAddress).RegisterRoutes(chipmunkRouter)
}
