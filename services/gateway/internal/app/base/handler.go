package base

import (
	gorilla "github.com/gorilla/mux"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	logger *log.Logger
}

var (
	controller *Controller
)

func ControllerInstance(logger *log.Logger) *Controller {
	if controller == nil {
		controller = &Controller{
			logger: logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	app := router.PathPrefix("/app").Subrouter()

	app.HandleFunc("/platforms", c.platforms).Methods(http.MethodGet)
}

type Platform struct {
	Name string
}

func (c Controller) platforms(res http.ResponseWriter, req *http.Request) {
	response := make([]Platform, 0)
	for _, platform := range api.Platform_name {
		d := Platform{
			Name: platform,
		}
		response = append(response, d)
	}
	httpext.SendModel(res, req, http.StatusOK, response)
}
