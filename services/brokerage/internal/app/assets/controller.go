package assets

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
	assetsService brokerageApi.AssetServiceClient
}

var (
	httpController *Controller
)

func NewController(connection grpc.ClientConnInterface) *Controller {
	if httpController == nil {
		httpController = &Controller{assetsService: brokerageApi.NewAssetServiceClient(connection)}
	}
	return httpController
}

func (c *Controller) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/assets", c.Set).Methods(http.MethodPut)
	router.HandleFunc("/assets/list/{page}", c.List).Methods(http.MethodGet)
	router.HandleFunc("/assets/{name}", c.Get).Methods(http.MethodGet)
}

func (c *Controller) Set(res http.ResponseWriter, req *http.Request) {
	model := new(api.Asset)
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
	}
	if _, err := c.assetsService.Set(req.Context(), model); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendCode(res, req, http.StatusOK)
	}
}

func (c *Controller) Get(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.GetAssetRequest)
	data.Name = muxext.PathParam(req, "name")
	if asset, err := c.assetsService.GetByName(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}

func (c Controller) List(res http.ResponseWriter, req *http.Request) {
	listReq := new(brokerageApi.GetAssetListRequest)
	page, err := strconv.Atoi(muxext.PathParam(req, "page"))
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	listReq.Page = int32(page)
	if list, err := c.assetsService.List(req.Context(), listReq); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, list)
	}
}
