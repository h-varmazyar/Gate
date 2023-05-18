package signals

import (
	gorilla "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	controller *Controller
)

type Controller struct {
	logger *log.Logger
}

func ControllerInstance(logger *log.Logger, eagleAddress string) *Controller {
	if controller == nil {
		controller = &Controller{
			logger: logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	_ = router.PathPrefix("/signals").Subrouter()
}
