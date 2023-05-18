package indicators

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	log "github.com/sirupsen/logrus"
)

var (
	controller *Controller
)

type Controller struct {
	indicatorsService chipmunkApi.IndicatorServiceClient
	logger            *log.Logger
}

func ControllerInstance(logger *log.Logger, chipmunkAddress string) *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(chipmunkAddress)
		controller = &Controller{
			indicatorsService: chipmunkApi.NewIndicatorServiceClient(chipmunkConn),
			logger:            logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	_ = router.PathPrefix("/indicators").Subrouter()
}
