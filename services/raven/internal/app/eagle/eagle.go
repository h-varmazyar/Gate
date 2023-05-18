package eagle

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/eagle/signals"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/eagle/strategies"
	log "github.com/sirupsen/logrus"
)

func RegisterRoutes(router *gorilla.Router, logger *log.Logger, configs *Configs) {
	eagleRouter := router.PathPrefix("/eagle").Subrouter()
	signals.ControllerInstance(logger, configs.EagleAddress).RegisterRoutes(eagleRouter)
	strategies.ControllerInstance(logger, configs.EagleAddress).RegisterRoutes(eagleRouter)
}
