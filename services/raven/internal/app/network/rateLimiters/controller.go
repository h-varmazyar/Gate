package rateLimiters

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	networkApi "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	rateLimitersService networkApi.RateLimiterServiceClient
	logger              *log.Logger
}

func ControllerInstance(logger *log.Logger, networkAddress string) *Controller {
	if controller == nil {
		networkConn := grpcext.NewConnection(networkAddress)
		controller = &Controller{
			rateLimitersService: networkApi.NewRateLimiterServiceClient(networkConn),
			logger:              logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	rateLimiter := router.PathPrefix("/rate-limiters").Subrouter()

	rateLimiter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	}).Methods(http.MethodPost)
}
