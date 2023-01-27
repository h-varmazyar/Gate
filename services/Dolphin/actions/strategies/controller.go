package strategies

//import (
//	"github.com/h-varmazyar/Gate/api/proto"
//	"github.com/h-varmazyar/Gate/pkg/errors"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	"github.com/h-varmazyar/Gate/services/Dolphin/configs"
//	"github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"
//	brokerageApi "github.com/h-varmazyar/Gate/services/core/api"
//	"net/http"
//)
//
//type strategyController struct {
//	strategyService brokerageApi.StrategyServiceClient
//}
//
//func newStrategyController() strategyController {
//	return strategyController{
//		strategyService: brokerageApi.NewStrategyServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
//	}
//}
//
//func (c *strategyController) create(ctx app.Context) error {
//
//	indicators := make([]*brokerageApi.Indicator, 0)
//	if _, err := c.strategyService.Create(ctx, &brokerageApi.CreateStrategyReq{
//		Name:        "",
//		Description: "",
//		Indicators:  indicators,
//	}); err != nil {
//		errModel := errors.Cast(ctx, err)
//		return ctx.Error(errModel.HttpStatus(), err)
//	}
//	return ctx.Redirect("/strategies/list")
//}
//
//func (c *strategyController) list(ctx app.Context) error {
//	if strategies, err := c.strategyService.List(ctx, new(proto.Void)); err != nil {
//		errModel := errors.Cast(ctx, err)
//		return ctx.Error(errModel.HttpStatus(), err)
//	} else {
//		ctx.Set("strategies", strategies.Elements)
//		return ctx.Render(http.StatusOK, "strategies/list")
//	}
//}
//
//func (c *strategyController) returnByID(ctx app.Context) error {
//	return nil
//}
