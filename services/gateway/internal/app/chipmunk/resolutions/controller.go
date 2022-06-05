package resolutions

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
	resolutionsService chipmunkApi.ResolutionServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		controller = &Controller{
			resolutionsService: chipmunkApi.NewResolutionServiceClient(chipmunkConn),
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *mux.Router) {
	resolutions := router.PathPrefix("/resolutions").Subrouter()
}
