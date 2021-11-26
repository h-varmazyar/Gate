package markets

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
	marketService brokerageApi.MarketServiceClient
}

var (
	httpController *Controller
)

func NewController(connection grpc.ClientConnInterface) *Controller {
	if httpController == nil {
		httpController = &Controller{marketService: brokerageApi.NewMarketServiceClient(connection)}
	}
	return httpController
}

func (c *Controller) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/markets", c.Set).Methods(http.MethodPut)
	router.HandleFunc("/markets/{id}", c.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/markets/brokerages/{brokerage_id}", c.List).Methods(http.MethodGet)
	router.HandleFunc("/markets/names/{name}/brokerages/{brokerage_id}", c.Get).Methods(http.MethodGet)
	router.HandleFunc("/markets/{id}/change-status", c.ChangeStatus).Methods(http.MethodGet)
}

func (c *Controller) Set(res http.ResponseWriter, req *http.Request) {
	model := new(brokerageApi.Market)
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if brokerage, err := c.marketService.Set(req.Context(), model); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, brokerage)
	}
}

func (c *Controller) List(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.MarketListRequest)
	data.BrokerageID = muxext.PathParam(req, "brokerage_id")
	if asset, err := c.marketService.List(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c *Controller) Get(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.MarketRequest)
	data.MarketName = muxext.PathParam(req, "name")
	data.BrokerageID = muxext.PathParam(req, "brokerage_id")
	if asset, err := c.marketService.Get(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c Controller) Delete(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusNotImplemented)
	//data := new(brokerageApi.BrokerageIDReq)
	//data.ID = muxext.PathParam(req, "id")
	//if list, err := c.brokerageService.Delete(req.Context(), data); err != nil {
	//	httpext.SendError(res, req, err)
	//} else {
	//	httpext.SendModel(res, req, http.StatusOK, list)
	//}
}

func (c Controller) ChangeStatus(res http.ResponseWriter, req *http.Request) {
	statusReq := new(api.StatusChangeRequest)
	statusReq.ID = muxext.PathParam(req, "id")
	if status, err := c.marketService.ChangeStatus(req.Context(), statusReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, map[string]string{"status": status.Status.String()})
	}
}
