package resolutions

import (
	"github.com/mrNobody95/Gate/pkg/grpcext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/dolphin/actions/viewHelpers"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"net/http"
	"strconv"
)

type resolutionController struct {
	resolutionService brokerageApi.ResolutionServiceClient
}

func newResolutionController() resolutionController {
	return resolutionController{
		resolutionService: brokerageApi.NewResolutionServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
	}
}

func (c *resolutionController) add(ctx app.Context) error {
	duration, err := strconv.ParseInt(ctx.Request().Form.Get("duration"), 10, 64)
	if err != nil {
		return ctx.Error(http.StatusNotAcceptable, err)
	}
	if _, err := c.resolutionService.Set(ctx, &brokerageApi.Resolution{
		BrokerageName: ctx.Request().Form.Get("brokerageRadio"),
		Duration:      duration,
		Label:         ctx.Request().Form.Get("label"),
		Value:         ctx.Request().Form.Get("value"),
	}); err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	return ctx.Redirect("/resolutions/list")
}

func (c *resolutionController) list(ctx app.Context) error {
	resolutions, err := c.resolutionService.List(ctx, &brokerageApi.GetResolutionListRequest{BrokerageName: brokerageApi.Names_All.String()})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("resolutions", resolutions.Resolutions)
	return ctx.Render(http.StatusOK, "resolutions/list", viewHelpers.Sum, viewHelpers.ResolutionLabel, viewHelpers.TimeStampFormat)
}

func (c *resolutionController) showBrokerageResolutions(ctx app.Context) error {
	name := ctx.Param("brokerage_name")
	resolutions, err := c.resolutionService.List(ctx, &brokerageApi.GetResolutionListRequest{BrokerageName: name})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}
	ctx.Set("resolutions", resolutions.Resolutions)
	return ctx.Render(http.StatusOK, "brokerages/resolution-table", viewHelpers.Sum, viewHelpers.ResolutionLabel)
}
