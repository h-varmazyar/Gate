package resolutions

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"github.com/h-varmazyar/gopack/mux"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	resolutionsService chipmunkApi.ResolutionServiceClient
}

func ControllerInstance() *Controller {
	if controller == nil {
		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		controller = &Controller{
			resolutionsService: chipmunkApi.NewResolutionServiceClient(chipmunkConn),
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
	listReq := new(chipmunkApi.GetResolutionListRequest)

	brokerageNames := mux.QueryParam(req, "brokerage-id")
	if len(brokerageNames) != 0 {
		listReq.BrokerageName = brokerageNames[0]
	}

	if resolutions, err := c.resolutionsService.List(req.Context(), listReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, resolutions.Elements)
	}
}

func (c Controller) get(res http.ResponseWriter, req *http.Request) {
	getReq := new(chipmunkApi.GetResolutionByIDRequest)

	if err := httpext.BindModel(req, getReq); err != nil {
		httpext.SendError(res, req, err)
		return
	}

	if resolution, err := c.resolutionsService.GetByID(req.Context(), getReq); err != nil {
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
