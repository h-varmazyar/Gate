package strategies

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	log "github.com/sirupsen/logrus"
)

var (
	controller *Controller
)

type Controller struct {
	strategiesService eagleApi.StrategyServiceClient
	logger            *log.Logger
}

func ControllerInstance(logger *log.Logger, eagleAddress string) *Controller {
	if controller == nil {
		eagleConn := grpcext.NewConnection(eagleAddress)
		controller = &Controller{
			strategiesService: eagleApi.NewStrategyServiceClient(eagleConn),
			logger:            logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	_ = router.PathPrefix("/strategies").Subrouter()
}
