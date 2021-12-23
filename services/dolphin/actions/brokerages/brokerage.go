package brokerages

import (
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/dolphin/actions/viewHelpers"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"net/http"
)

type brokerageController struct {
	brokerageService  brokerageApi.BrokerageServiceClient
	marketService     brokerageApi.MarketServiceClient
	resolutionService brokerageApi.ResolutionServiceClient
}

func newBrokerageController() brokerageController {
	return brokerageController{
		brokerageService:  brokerageApi.NewBrokerageServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
		marketService:     brokerageApi.NewMarketServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
		resolutionService: brokerageApi.NewResolutionServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
	}
}

func (c *brokerageController) list(ctx app.Context) error {
	brokerages, err := c.brokerageService.List(ctx, new(api.Void))
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("brokerages", brokerages.Brokerages)
	return ctx.Render(http.StatusOK, "brokerages/list", viewHelpers.Sum, viewHelpers.ResolutionLabel)
}

func (c *brokerageController) show(ctx app.Context) error {
	brokerage, err := c.brokerageService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: ctx.Param("brokerage_id")})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	ctx.Set("brokerage", brokerage)
	return ctx.Render(http.StatusOK, "brokerages/show", viewHelpers.Sum, viewHelpers.ResolutionLabel)
}

func (c *brokerageController) overview(ctx app.Context) error {
	fmt.Println("in overview")
	return ctx.Render(http.StatusOK, "brokerages/overview")
}

func (c *brokerageController) showAddPage(ctx app.Context) error {
	ctx.Set("resolutions", make([]*brokerageApi.Resolution, 0))
	ctx.Set("markets", make([]*brokerageApi.Market, 0))
	return ctx.Render(http.StatusOK, "brokerages/add", viewHelpers.Sum, viewHelpers.ResolutionLabel)
}

func (c *brokerageController) add(ctx app.Context) error {
	//todo: add markets and resolutions from input
	brokerage := new(brokerageApi.CreateBrokerageReq)
	brokerage.Title = ctx.Request().Form.Get("br-title")
	brokerage.Description = ctx.Request().Form.Get("description")
	if brokerageName, ok := brokerageApi.Names_value[ctx.Request().Form.Get("brokerageRadio")]; ok {
		brokerage.Name = brokerageApi.Names(brokerageName)
	} else {
		fmt.Println("brokerage name not:", ctx.Request().Form.Get("brokerageRadio"))
	}
	brokerage.Auth = &api.Auth{
		AccessID:  ctx.Request().Form.Get("access-id"),
		SecretKey: ctx.Request().Form.Get("secret-key"),
	}
	brokerage.Markets = new(brokerageApi.Markets)
	for _, marketID := range ctx.Request().Form["marketsCheckbox"] {
		brokerage.Markets.Markets = append(brokerage.Markets.Markets, &brokerageApi.Market{ID: marketID})
	}
	brokerage.ResolutionID = ctx.Request().Form.Get("resolutionRadio")
	brokerage.StrategyID = ctx.Request().Form.Get("strategyRadio")
	_, err := c.brokerageService.Create(ctx, brokerage)
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	return ctx.Redirect("/brokerages/list")
}
