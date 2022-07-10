package brokerages

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"net/http"
)

var (
	handler *Handler
)

type Handler struct {
	brokerageService brokerageApi.BrokerageServiceClient
}

func HandlerInstance() *Handler {
	if handler == nil {
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		handler = &Handler{
			brokerageService: brokerageApi.NewBrokerageServiceClient(brokerageConn),
		}
	}
	return handler
}

func (h Handler) RegisterRoutes(router *gorilla.Router) {
	brokerage := router.PathPrefix("/brokerages").Subrouter()

	brokerage.HandleFunc("/create", h.create).Methods(http.MethodPost)
	brokerage.HandleFunc("/active", h.returnActive).Methods(http.MethodGet)
	brokerage.HandleFunc("/start", h.start).Methods(http.MethodPost)
	brokerage.HandleFunc("/stop", h.stop).Methods(http.MethodPost)
}

func (h Handler) create(res http.ResponseWriter, req *http.Request) {
	brokerage := new(brokerageApi.CreateBrokerageReq)
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
	if brokerage, err := h.brokerageService.Active(req.Context(), &brokerageApi.ActiveBrokerageReq{
		WithMarkets:     true,
		WithResolutions: true,
	}); err != nil {
		httpext.SendError(res, req, err)
	} else {
		response := new(Brokerage)
		mapper.Struct(brokerage, response)
		httpext.SendModel(res, req, http.StatusOK, response)
	}
}

func (h Handler) start(res http.ResponseWriter, req *http.Request) {
	start := new(Start)
	if err := httpext.BindModel(req, start); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	change := new(brokerageApi.StatusChangeRequest)
	mapper.Struct(start, change)

	if status, err := h.brokerageService.ChangeStatus(req.Context(), change); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, status)
	}
}

func (h Handler) stop(res http.ResponseWriter, req *http.Request) {
	change := new(brokerageApi.StatusChangeRequest)
	if status, err := h.brokerageService.ChangeStatus(req.Context(), change); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, status)
	}
}
