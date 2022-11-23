package functions

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
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
		handler = &Handler{
			functionsService: coreApi.NewFunctionsServiceClient(coreConn),
			logger:           logger,
		}
	}
	return handler
}

func (h Handler) RegisterRoutes(router *gorilla.Router) {
}
