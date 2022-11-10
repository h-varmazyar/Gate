package functions

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"net/http"
)

var (
	handler *Handler
)

type Handler struct {
	functionsService coreApi.FunctionsServiceClient
}

func HandlerInstance() *Handler {
	if handler == nil {
		coreConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Core)
		handler = &Handler{
			functionsService: coreApi.NewFunctionsServiceClient(coreConn),
		}
	}
	return handler
}

func (h Handler) RegisterRoutes(router *gorilla.Router) {
	brokerage := router.PathPrefix("/brokerages").Subrouter()

	brokerage.HandleFunc("/testAsync", h.testAsync).Methods(http.MethodPost)
}

func (h Handler) testAsync(res http.ResponseWriter, req *http.Request) {
	if _, err := h.functionsService.AsyncTest(req.Context(), new(api.Void)); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusCreated)
	}
}
