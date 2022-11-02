package markets

import (
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/Dolphin/actions/viewHelpers"
	"github.com/h-varmazyar/Gate/services/Dolphin/configs"
	"github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"
	brokerageApi "github.com/h-varmazyar/Gate/services/core/api"
	"net/http"
)

type marketController struct {
	marketService brokerageApi.MarketServiceClient
}

func newMarketController() marketController {
	return marketController{
		marketService: brokerageApi.NewMarketServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
	}
}

func (c *marketController) add(ctx app.Context) error {
	req := new(brokerageApi.UpdateMarketsReq)
	req.BrokerageName = ctx.Request().Form.Get("brokerageRadio")
	_, err := c.marketService.UpdateMarkets(ctx, req)
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	return ctx.Redirect("/markets/list")
}

func (c *marketController) show(ctx app.Context) error {
	fmt.Println("show markets")
	return ctx.Render(http.StatusOK, "/markets/show")
}

func (c *marketController) list(ctx app.Context) error {
	markets, err := c.marketService.List(ctx, &brokerageApi.MarketListRequest{BrokerageName: brokerageApi.Names_All.String()})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("markets", markets.Markets)
	return ctx.Render(http.StatusOK, "/markets/list", viewHelpers.Sum, viewHelpers.TimeStampFormat)
}

func (c *marketController) showBrokerageMarkets(ctx app.Context) error {
	name := ctx.Param("brokerage_name")
	markets, err := c.marketService.List(ctx, &brokerageApi.MarketListRequest{BrokerageName: name})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("markets", markets.Markets)
	return ctx.Render(http.StatusOK, "brokerages/markets-table", viewHelpers.Sum)
}
