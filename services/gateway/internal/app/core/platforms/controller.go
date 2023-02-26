package platforms

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
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
	_ = router.PathPrefix("/platforms").Subrouter()
}
