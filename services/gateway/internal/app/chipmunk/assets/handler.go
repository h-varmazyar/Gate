package assets

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
)

var (
	controller *Controller
)

type Controller struct {
	assetsService chipmunkApi.AssetServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		controller = &Controller{
			assetsService: chipmunkApi.NewAssetServiceClient(chipmunkConn),
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	_ = router.PathPrefix("/assets").Subrouter()
}
