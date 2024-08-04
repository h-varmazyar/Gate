package resolutions

import (
	gorilla "github.com/gorilla/mux"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	resolutionsService chipmunkApi.ResolutionServiceClient
	logger             *log.Logger
}

func ControllerInstance(logger *log.Logger, chipmunkAddress string) *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(chipmunkAddress)
		controller = &Controller{
			resolutionsService: chipmunkApi.NewResolutionServiceClient(chipmunkConn),
			logger:             logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gorilla.Router) {
	resolutions := router.PathPrefix("/resolutions").Subrouter()

	resolutions.HandleFunc("/create", c.create).Methods(http.MethodPost)
	resolutions.HandleFunc("/list", c.list).Methods(http.MethodGet)
	resolutions.HandleFunc("/{resolution-id}", c.get).Methods(http.MethodGet)
	resolutions.HandleFunc("/{resolution-id}", c.update).Methods(http.MethodPut)
}

// resolutionCreate godoc
//
// @Summary     Create new resolution manually
// @Description Create new resolution manually
// @Accept      json
// @Produce     json
// @Param       resolution body ResolutionSetReq true "New Resolution"
// @Success     201
// @Failure     400 {object} errors.Error
// @Failure     404 {object} errors.Error
// @Failure     500 {object} errors.Error
// @Router      /chipmunk/resolutions/create [post]
func (c Controller) create(res http.ResponseWriter, req *http.Request) {
	resolution := new(chipmunkApi.Resolution)
	if err := httpext.BindModel(req, resolution); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.resolutionsService.Set(req.Context(), resolution); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusCreated)
	}
}

// resolutionList godoc
//
// @Summary     get resolution list
// @Description get resolution list based on platform
// @Accept      json
// @Produce     json
// @Param       platform query  string   true "Platform name" Enums:(Coinex,UnknownBrokerage,Nobitex,Mazdax,Binance)
// @Success     200             {object} proto.Resolutions
// @Failure     400             {object} errors.Error
// @Failure     404             {object} errors.Error
// @Failure     500             {object} errors.Error
// @Router      /chipmunk/resolutions/list [get]
func (c Controller) list(res http.ResponseWriter, req *http.Request) {
	listReq := new(chipmunkApi.ResolutionListReq)

	platforms := mux.QueryParam(req, "platform")
	if len(platforms) == 0 {
		httpext.SendError(res, req, errors.New(req.Context(), codes.InvalidArgument).AddDetails("platform needed"))
		return
	}
	if platforms[0] != "" {
		listReq.Platform = api.Platform(api.Platform_value[platforms[0]])
	}

	if resolutions, err := c.resolutionsService.List(req.Context(), listReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, resolutions.Elements)
	}
}

// resolutionList godoc
//
// @Summary     get single resolution
// @Description get single resolution based on resolution id
// @Accept      json
// @Produce     json
// @Param       resolution-id path  string true     "Resolution id"
// @Success     200                        {object} proto.Resolution
// @Failure     400                        {object} errors.Error
// @Failure     404                        {object} errors.Error
// @Failure     500                        {object} errors.Error
// @Router      /chipmunk/resolutions/{resolution-id} [get]
func (c Controller) get(res http.ResponseWriter, req *http.Request) {
	getReq := new(chipmunkApi.ResolutionReturnByIDReq)

	getReq.ID = mux.PathParam(req, "resolution-id")

	if resolution, err := c.resolutionsService.ReturnByID(req.Context(), getReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, resolution)
	}
}

// resolutionList godoc
//
// @Summary     update single resolution
// @Description update single resolution based on resolution id
// @Accept      json
// @Produce     json
// @Param       resolution-id path string                        true "Resolution id"
// @Param       resolution         body   ResolutionSetReq true "New Resolution"
// @Success     200
// @Failure     400 {object} errors.Error
// @Failure     404 {object} errors.Error
// @Failure     500 {object} errors.Error
// @Router      /chipmunk/resolutions/{resolution-id} [put]
func (c Controller) update(res http.ResponseWriter, req *http.Request) {
	resolution := new(chipmunkApi.Resolution)
	if err := httpext.BindModel(req, resolution); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	resolution.ID = mux.PathParam(req, "resolution-id")
	if _, err := c.resolutionsService.Set(req.Context(), resolution); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}
