package resolutions

import (
	gorilla "github.com/gorilla/mux"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
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
	resolutions.HandleFunc("/{resolution-id}", c.delete).Methods(http.MethodDelete)
}

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

func (c Controller) list(res http.ResponseWriter, req *http.Request) {
	listReq := new(chipmunkApi.ResolutionListReq)

	platforms := mux.QueryParam(req, "platform")
	if len(platforms) != 0 {
		listReq.Platform = api.Platform(api.Platform_value[platforms[0]])
	}

	if resolutions, err := c.resolutionsService.List(req.Context(), listReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, resolutions.Elements)
	}
}

func (c Controller) get(res http.ResponseWriter, req *http.Request) {
	getReq := new(chipmunkApi.ResolutionReturnByIDReq)

	if err := httpext.BindModel(req, getReq); err != nil {
		httpext.SendError(res, req, err)
		return
	}

	if resolution, err := c.resolutionsService.ReturnByID(req.Context(), getReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, resolution)
	}
}

func (c Controller) update(res http.ResponseWriter, req *http.Request) {
	resolution := new(chipmunkApi.Resolution)
	if err := httpext.BindModel(req, resolution); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.resolutionsService.Set(req.Context(), resolution); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}

func (c Controller) delete(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusNotFound)
}
