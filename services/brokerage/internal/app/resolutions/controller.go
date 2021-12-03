package resolutions

import (
	"github.com/gorilla/mux"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/httpext"
	"github.com/mrNobody95/Gate/pkg/muxext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
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
	ResolutionsService brokerageApi.ResolutionServiceClient
}

var (
	httpController *Controller
)

func NewController(connection grpc.ClientConnInterface) *Controller {
	if httpController == nil {
		httpController = &Controller{ResolutionsService: brokerageApi.NewResolutionServiceClient(connection)}
	}
	return httpController
}

func (c *Controller) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/resolutions", c.Set).Methods(http.MethodPost)
	router.HandleFunc("/resolutions/{resolution_id}", c.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/resolutions/brokerages/{brokerage_id}", c.List).Methods(http.MethodGet)
	router.HandleFunc("/resolutions/{resolution_duration}/brokerages/{brokerage_id}", c.GetByDuration).Methods(http.MethodGet)
}

func (c *Controller) Set(res http.ResponseWriter, req *http.Request) {
	model := new(api.Resolution)
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	if _, err := c.ResolutionsService.Set(req.Context(), model); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusCreated)
	}
}

func (c *Controller) List(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.GetResolutionListRequest)
	data.BrokerageID = muxext.PathParam(req, "brokerage_id")
	if asset, err := c.ResolutionsService.List(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c *Controller) GetByID(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.GetResolutionByIDRequest)
	data.ID = muxext.PathParam(req, "resolution_id")
	if asset, err := c.ResolutionsService.GetByID(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c *Controller) GetByDuration(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.GetResolutionByDurationRequest)
	data.BrokerageID = muxext.PathParam(req, "brokerage_id")
	duration, err := strconv.ParseInt(muxext.PathParam(req, "resolution_duration"), 10, 64)
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	data.Duration = duration
	if asset, err := c.ResolutionsService.GetByDuration(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}
