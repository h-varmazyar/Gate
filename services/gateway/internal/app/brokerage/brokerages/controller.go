package brokerages

import (
	"github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
)

var (
	controller *Controller
)

type Controller struct {
	brokerageService brokerageApi.BrokerageServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		controller = &Controller{
			brokerageService: brokerageApi.NewBrokerageServiceClient(brokerageConn),
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *mux.Router) {
	brokerage := router.PathPrefix("/brokerages").Subrouter()
}
