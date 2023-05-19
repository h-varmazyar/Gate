package candles

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
	candlesService chipmunkApi.CandleServiceClient
	logger         *log.Logger
}

func ControllerInstance(logger *log.Logger, chipmunkAddress string) *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(chipmunkAddress)
		controller = &Controller{
			candlesService: chipmunkApi.NewCandleServiceClient(chipmunkConn),
			logger:         logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	_ = router.PathPrefix("/rateLimiters").Subrouter()

}
