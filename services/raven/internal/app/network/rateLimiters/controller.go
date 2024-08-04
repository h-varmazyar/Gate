package rateLimiters

import (
	"fmt"
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	networkApi "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/gopack/mux"
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

	rateLimiter.HandleFunc("/create", c.create).Methods(http.MethodPost)
	rateLimiter.HandleFunc("/list", c.list).Methods(http.MethodGet)
	rateLimiter.HandleFunc("/{rate_limiter_id}", c.returnByID).Methods(http.MethodGet)
}

// rateLimiterCreate godoc
//
// @Summary     Create new rate limiter
// @Description Create new rate limiter
// @Accept      json
// @Produce     json
// @Param       RateLimiter body  RateLimiter true "New rate limiter"
// @Success     201               {object}    proto.RateLimiter
// @Failure     400               {object}    errors.Error
// @Failure     404               {object}    errors.Error
// @Failure     500               {object}    errors.Error
// @Router      /network/rate-limiters/create [post]
func (c Controller) create(res http.ResponseWriter, req *http.Request) {
	rateLimiterCreateReq := new(networkApi.RateLimiterCreateReq)
	if err := httpext.BindModel(req, rateLimiterCreateReq); err != nil {
		httpext.SendError(res, req, err)
		return
	}

	if ip, err := c.rateLimitersService.Create(req.Context(), rateLimiterCreateReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusCreated, ip)
	}
}

// rateLimiterList godoc
//
// @Summary     return rate limiter list
// @Description return rate limiter list
// @Accept      json
// @Produce     json
// @Param       type query          string true "rate limiter type" Enums:(Spread,Immediate)
// @Success     200        {object} proto.RateLimiters
// @Failure     400        {object} errors.Error
// @Failure     404        {object} errors.Error
// @Failure     500        {object} errors.Error
// @Router      /network/rate-limiters/list [get]
func (c Controller) list(res http.ResponseWriter, req *http.Request) {
	if rateLimiters, err := c.rateLimitersService.List(req.Context(), new(networkApi.RateLimiterListReq)); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, rateLimiters)
	}
}

// rateLimiterReturn godoc
//
// @Summary     return rate limiter with ID
// @Description return rate limiter with ID
// @Accept      json
// @Produce     json
// @Param       rate_limiter_id path  string true     "rate limiter ID"
// @Success     200                          {object} proto.RateLimiter
// @Failure     400                          {object} errors.Error
// @Failure     404                          {object} errors.Error
// @Failure     500                          {object} errors.Error
// @Router      /network/rate-limiters/{rate_limiter_id} [get]
func (c Controller) returnByID(res http.ResponseWriter, req *http.Request) {
	rateLimiterReturnReq := new(networkApi.RateLimiterReturnReq)

	rateLimiterReturnReq.ID = mux.PathParam(req, "rate_limiter_id")

	fmt.Println("id is:", rateLimiterReturnReq.ID)

	if rateLimiter, err := c.rateLimitersService.Return(req.Context(), rateLimiterReturnReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, rateLimiter)
	}
}
