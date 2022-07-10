package signals

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
)

var (
	controller *Controller
)

type Controller struct {
	signalsService eagleApi.SignalServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		eagleConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		controller = &Controller{
			signalsService: eagleApi.NewSignalServiceClient(eagleConn),
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	signals := router.PathPrefix("/signals").Subrouter()
}
