package brokerages

import (
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/Dolphin/actions/viewHelpers"
	"github.com/h-varmazyar/Gate/services/Dolphin/configs"
	"github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"
	brokerageApi "github.com/h-varmazyar/Gate/services/core/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/http"
	"strconv"
)

type brokerageController struct {
	brokerageService  brokerageApi.BrokerageServiceClient
	marketService     brokerageApi.MarketServiceClient
	strategyService   brokerageApi.StrategyServiceClient
	resolutionService brokerageApi.ResolutionServiceClient
}

func newBrokerageController() brokerageController {
	return brokerageController{
		brokerageService:  brokerageApi.NewBrokerageServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
		marketService:     brokerageApi.NewMarketServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
		strategyService:   brokerageApi.NewStrategyServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
		resolutionService: brokerageApi.NewResolutionServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
	}
}

func (c *brokerageController) list(ctx app.Context) error {
	brokerages, err := c.brokerageService.List(ctx, new(api.Void))
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("brokerages", brokerages.Brokerages)
	ctx.Set("enable", api.Status_Enable)
	ctx.Set("disable", api.Status_Disable)
	return ctx.Render(http.StatusOK, "brokerages/list", viewHelpers.Sum, viewHelpers.ResolutionLabel)
}

func (c *brokerageController) show(ctx app.Context) error {
	id, err := strconv.ParseUint(ctx.Param("brokerage_id"), 10, 32)
	if err != nil {
		return err
	}
	brokerage, err := c.brokerageService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: uint32(id)})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	ctx.Set("brokerage", brokerage)
	return ctx.Render(http.StatusOK, "brokerages/show", viewHelpers.Sum, viewHelpers.ResolutionLabel)
}

func (c *brokerageController) overview(ctx app.Context) error {
	return ctx.Render(http.StatusOK, "brokerages/overview")
}

func (c *brokerageController) switchStatus(ctx app.Context) error {
	ohlc := false
	trading := false
	if ctx.Request().Form.Get("ohlcCheckbox") != "" {
		ohlc = true
	}
	if ctx.Request().Form.Get("tradingCheckbox") != "" {
		trading = true
	}
	id, err := strconv.ParseUint(ctx.Param("brokerage_id"), 10, 32)
	if err != nil {
		return err
	}
	if _, err := c.brokerageService.ChangeStatus(ctx, &brokerageApi.StatusChangeRequest{
		ID:      uint32(id),
		OHLC:    ohlc,
		Trading: trading,
	}); err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	return ctx.Redirect("/brokerages/list")
}

func (c *brokerageController) showAddPage(ctx app.Context) error {
	strategies, err := c.strategyService.List(ctx, new(api.Void))
	if err != nil {
		log.WithError(err).Error("failed to load strategies")
		return errors.NewWithSlug(ctx, codes.FailedPrecondition, "failed to load strategiess")
	}
	ctx.Set("resolutions", make([]*brokerageApi.Resolution, 0))
	ctx.Set("markets", make([]*brokerageApi.Market, 0))
	ctx.Set("strategies", strategies.Elements)
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
		log.Error("brokerage name not:", ctx.Request().Form.Get("brokerageRadio"))
		return errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong brokerage name")
	}
	brokerage.Auth = &api.Auth{
		AccessID:  ctx.Request().Form.Get("access-id"),
		SecretKey: ctx.Request().Form.Get("secret-key"),
	}
	brokerage.Markets = new(brokerageApi.Markets)
	for _, marketID := range ctx.Request().Form["marketsCheckbox"] {
		id, err := strconv.ParseUint(marketID, 10, 32)
		if err != nil {
			return err
		}
		brokerage.Markets.Markets = append(brokerage.Markets.Markets, &brokerageApi.Market{ID: uint32(id)})
	}
	brokerage.ResolutionID = ctx.Request().Form.Get("resolutionRadio")
	strategyID, err := strconv.ParseUint(ctx.Request().Form.Get("strategyRadio"), 10, 32)
	if err != nil {
		return err
	}
	brokerage.StrategyID = uint32(strategyID)
	_, err = c.brokerageService.Create(ctx, brokerage)
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	return ctx.Redirect("/brokerages/list")
}
