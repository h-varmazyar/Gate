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

	candles.HandleFunc("/primary-candles", c.downloadPrimaryCandles).Methods(http.MethodPost)
}

func (c Controller) downloadPrimaryCandles(res http.ResponseWriter, req *http.Request) {
	worker := new(chipmunkApi.DownloadPrimaryCandlesReq)
	if err := httpext.BindModel(req, worker); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.candlesService.DownloadPrimaryCandles(req.Context(), worker); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}
