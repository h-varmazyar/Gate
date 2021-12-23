package markets

import (
	"github.com/mrNobody95/Gate/pkg/grpcext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/dolphin/actions/viewHelpers"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
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

func (c *marketController) showBrokerageMarkets(ctx app.Context) error {
	name := ctx.Param("brokerage_name")
	markets, err := c.marketService.List(ctx, &brokerageApi.MarketListRequest{BrokerageName: name})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("markets", markets.Markets)
	return ctx.Render(http.StatusOK, "brokerages/markets-table", viewHelpers.Sum)
}
