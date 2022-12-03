package platforms

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	handler *Controller
)

type Controller struct {
	platformService coreApi.PlatformServiceClient
	logger          *log.Logger
}

func HandlerInstance(logger *log.Logger, coreAddress string) *Controller {
	if handler == nil {
		coreConn := grpcext.NewConnection(coreAddress)
		handler = &Controller{
			platformService: coreApi.NewPlatformServiceClient(coreConn),
			logger:          logger,
		}
	}
	return handler
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	brokerage := router.PathPrefix("/platforms").Subrouter()

	brokerage.HandleFunc("/collect-data", c.collectMarketData).Methods(http.MethodPost)
}

func (c Controller) collectMarketData(res http.ResponseWriter, req *http.Request) {
	platformDataCollection := new(coreApi.PlatformCollectDataReq)
	if err := httpext.BindModel(req, platformDataCollection); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	_, err := c.platformService.CollectMarketData(req.Context(), platformDataCollection)
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	httpext.SendCode(res, req, http.StatusOK)
}
