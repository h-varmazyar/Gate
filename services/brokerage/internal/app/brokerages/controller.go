package assets

import (
	"github.com/gorilla/mux"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/httpext"
	"github.com/mrNobody95/Gate/pkg/muxext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
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
	brokerageService brokerageApi.BrokerageServiceClient
}

var (
	httpController *Controller
)

func NewController(connection grpc.ClientConnInterface) *Controller {
	if httpController == nil {
		httpController = &Controller{brokerageService: brokerageApi.NewBrokerageServiceClient(connection)}
	}
	return httpController
}

func (c *Controller) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/brokerages", c.Add).Methods(http.MethodPost)
	router.HandleFunc("/brokerages/{id}", c.Get).Methods(http.MethodGet)
	router.HandleFunc("/brokerages/{id}", c.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/brokerages/{id}/change-status", c.ChangeStatus).Methods(http.MethodGet)
}

func (c *Controller) Add(res http.ResponseWriter, req *http.Request) {
	model := new(brokerageApi.Brokerage)
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if brokerage, err := c.brokerageService.Add(req.Context(), model); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusCreated, brokerage)
	}
}

func (c *Controller) Get(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.BrokerageIDReq)
	data.ID = muxext.PathParam(req, "id")
	if asset, err := c.brokerageService.Get(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c Controller) Delete(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.BrokerageIDReq)
	data.ID = muxext.PathParam(req, "id")
	if list, err := c.brokerageService.Delete(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, list)
	}
}

func (c Controller) ChangeStatus(res http.ResponseWriter, req *http.Request) {
	statusReq := new(api.StatusChangeRequest)
	statusReq.ID = muxext.PathParam(req, "id")
	if status, err := c.brokerageService.ChangeStatus(req.Context(), statusReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, map[string]string{"status": status.Status.String()})
	}
}
