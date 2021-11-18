package app

import (
	"github.com/gorilla/mux"
	"github.com/mrNobody95/Gate/pkg/httpext"
	"github.com/mrNobody95/Gate/pkg/muxext"
	ganjehAPI "github.com/mrNobody95/Gate/services/ganjeh/api"
	"google.golang.org/grpc"
	"net/http"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 13.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Controller struct {
	ganjehService ganjehAPI.VariableServiceClient
}

var (
	httpController *Controller
)

func NewController(connection grpc.ClientConnInterface) *Controller {
	if httpController == nil {
		httpController = &Controller{ganjehService: ganjehAPI.NewVariableServiceClient(connection)}
	}
	return httpController
}

func (c *Controller) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/variables", c.Set).Methods(http.MethodPut)
	router.HandleFunc("/variables/namespace/{namespace}", c.List).Methods(http.MethodGet)
	router.HandleFunc("/variables/namespace/{namespace}/key/{key}", c.Get).Methods(http.MethodGet)
	router.HandleFunc("/variables/namespace/{namespace}/key/{key}", c.Delete).Methods(http.MethodDelete)
}

func (c *Controller) Set(res http.ResponseWriter, req *http.Request) {
	model := new(ganjehAPI.SetVariableRequest)
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
	}
	if _, err := c.ganjehService.Set(req.Context(), model); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}

func (c *Controller) Get(res http.ResponseWriter, req *http.Request) {
	data := new(ganjehAPI.GetVariableRequest)
	data.Key = muxext.PathParam(req, "key")
	data.Namespace = muxext.PathParam(req, "namespace")
	if asset, err := c.ganjehService.Get(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c Controller) List(res http.ResponseWriter, req *http.Request) {
	variablesReq := new(ganjehAPI.GetVariablesRequest)
	variablesReq.Namespace = muxext.PathParam(req, "namespace")
	if list, err := c.ganjehService.List(req.Context(), variablesReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, list)
	}
}

func (c Controller) Delete(res http.ResponseWriter, req *http.Request) {
	deleteReq := new(ganjehAPI.DeleteVariableRequest)
	deleteReq.Namespace = muxext.PathParam(req, "namespace")
	deleteReq.Key = muxext.PathParam(req, "key")
	if list, err := c.ganjehService.Delete(req.Context(), deleteReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusNoContent, list)
	}
}
