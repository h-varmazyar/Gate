package markets

import (
	"github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
)

var (
	controller *Controller
)

type Controller struct {
	marketsService chipmunkApi.MarketServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		controller = &Controller{
			marketsService: chipmunkApi.NewMarketServiceClient(chipmunkConn),
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *mux.Router) {
	markets := router.PathPrefix("/markets").Subrouter()
}
