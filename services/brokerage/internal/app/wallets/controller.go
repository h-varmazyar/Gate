package assets

import (
	"github.com/gorilla/mux"
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
	walletService brokerageApi.WalletServiceClient
}

var (
	httpController *Controller
)

func NewController(connection grpc.ClientConnInterface) *Controller {
	if httpController == nil {
		httpController = &Controller{walletService: brokerageApi.NewWalletServiceClient(connection)}
	}
	return httpController
}

func (c *Controller) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/wallets/brokerages/{brokerage_id}/list", c.List).Methods(http.MethodGet)
}

func (c *Controller) List(res http.ResponseWriter, req *http.Request) {
	data := new(brokerageApi.WalletListRequest)
	data.BrokerageID = muxext.PathParam(req, "brokerage_id")
	if asset, err := c.walletService.List(req.Context(), data); err != nil {
		httpext.SendError(res, req, err)
	} else {
		httpext.SendModel(res, req, http.StatusOK, asset)
	}
}
