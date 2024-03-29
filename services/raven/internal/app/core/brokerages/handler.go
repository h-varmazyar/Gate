package brokerages

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	handler *Handler
)

type Handler struct {
	brokerageService coreApi.BrokerageServiceClient
	logger           *log.Logger
}

func HandlerInstance(logger *log.Logger, coreAddress string) *Handler {
	if handler == nil {
		coreConn := grpcext.NewConnection(coreAddress)
		handler = &Handler{
			brokerageService: coreApi.NewBrokerageServiceClient(coreConn),
			logger:           logger,
		}
	}
	return handler
}

func (h Handler) RegisterRoutes(router *gorilla.Router) {
	brokerage := router.PathPrefix("/brokerages").Subrouter()

	brokerage.HandleFunc("/create", h.create).Methods(http.MethodPost)
	//core.HandleFunc("/active", h.returnActive).Methods(http.MethodGet)
	//core.HandleFunc("/start", h.start).Methods(http.MethodPost)
	//core.HandleFunc("/stop", h.stop).Methods(http.MethodPost)
}

func (h Handler) create(res http.ResponseWriter, req *http.Request) {
	brokerage := new(coreApi.BrokerageCreateReq)
	if err := httpext.BindModel(req, brokerage); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := h.brokerageService.Create(req.Context(), brokerage); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusCreated)
	}
}

func (h Handler) returnActive(res http.ResponseWriter, req *http.Request) {
	//if core, err := h.brokerageService.Active(req.Context(), &coreApi.ActiveBrokerageReq{
	//	WithMarkets:     true,
	//	WithResolutions: true,
	//}); err != nil {
	//	httpext.SendError(res, req, err)
	//} else {
	//	response := new(Brokerage)
	//	mapper.Struct(core, response)
	//	httpext.SendModel(res, req, http.StatusOK, response)
	//}
}

func (h Handler) start(res http.ResponseWriter, req *http.Request) {
	//start := new(Start)
	//if err := httpext.BindModel(req, start); err != nil {
	//	httpext.SendError(res, req, err)
	//	return
	//}
	//change := new(coreApi.StatusChangeRequest)
	//mapper.Struct(start, change)
	//
	//if status, err := h.brokerageService.Start(req.Context(), change); err != nil {
	//	httpext.SendError(res, req, err)
	//} else {
	//	httpext.SendModel(res, req, http.StatusOK, status)
	//}
}

func (h Handler) stop(res http.ResponseWriter, req *http.Request) {
	//change := new(coreApi.StatusChangeRequest)
	//if status, err := h.brokerageService.Stop(req.Context(), change); err != nil {
	//	httpext.SendError(res, req, err)
	//} else {
	//	httpext.SendModel(res, req, http.StatusOK, status)
	//}
}
