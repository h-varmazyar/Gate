package functions

import (
	gorilla "github.com/gorilla/mux"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	muxext "github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	handler *Handler
)

type Handler struct {
	functionsService coreApi.FunctionsServiceClient
	logger           *log.Logger
}

func HandlerInstance(logger *log.Logger, coreAddress string) *Handler {
	if handler == nil {
		coreConn := grpcext.NewConnection(coreAddress)
		logger.Infof("core connection is: %v", coreAddress)
		handler = &Handler{
			functionsService: coreApi.NewFunctionsServiceClient(coreConn),
			logger:           logger,
		}
	}
	return handler
}

func (h *Handler) RegisterRoutes(router *gorilla.Router) {
	functions := router.PathPrefix("/functions").Subrouter()

	functions.HandleFunc("/platform/{platform}/market/{market_name}", h.marketStatus).Methods(http.MethodGet)
}

func (h *Handler) marketStatus(res http.ResponseWriter, req *http.Request) {
	marketName := muxext.PathParam(req, "market_name")
	platform := muxext.PathParam(req, "platform")
	h.logger.Infof("market status called for: %v - %v", platform, marketName)
	marketStat, err := h.functionsService.SingleMarketStatistics(req.Context(), &coreApi.MarketStatisticsReq{
		MarketName: marketName,
		Platform:   api.Platform(api.Platform_value[platform]),
	})
	if err != nil {
		h.logger.WithError(err).Errorf("failed to get market status")
		httpext.SendError(res, req, err)
		return
	}

	h.logger.Infof("market status id: %v", marketStat)
	httpext.SendModel(res, req, http.StatusOK, marketStat)
}
