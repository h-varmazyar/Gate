package strategies

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	strategies := router.PathPrefix("/strategies").Subrouter()
	strategies.HandleFunc("/StartWorker", c.startWorker).Methods(http.MethodPost)
	strategies.HandleFunc("/StopWorker", c.stopWorker).Methods(http.MethodPost)
}

func (c Controller) startWorker(res http.ResponseWriter, req *http.Request) {
	worker := new(eagleApi.StrategySignalCheckStartReq)
	if err := httpext.BindModel(req, worker); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.strategiesService.StartSignalChecker(req.Context(), worker); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}

func (c Controller) stopWorker(res http.ResponseWriter, req *http.Request) {
	worker := new(eagleApi.StrategySignalCheckStopReq)
	if err := httpext.BindModel(req, worker); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.strategiesService.StopSignalChecker(req.Context(), worker); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}
