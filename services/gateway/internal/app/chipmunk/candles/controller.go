package candles

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	candles := router.PathPrefix("/candles").Subrouter()

	candles.HandleFunc("/workers/start", c.startWorkers).Methods(http.MethodPost)
	candles.HandleFunc("/workers/stop", c.stopWorkers).Methods(http.MethodPost)
}

func (c Controller) startWorkers(res http.ResponseWriter, req *http.Request) {
	worker := new(chipmunkApi.CandleWorkerStartReq)
	if err := httpext.BindModel(req, worker); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.candlesService.StartWorkers(req.Context(), worker); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}

func (c Controller) stopWorkers(res http.ResponseWriter, req *http.Request) {
	worker := new(chipmunkApi.CandleWorkerStopReq)
	if err := httpext.BindModel(req, worker); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.candlesService.StopWorkers(req.Context(), worker); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}
