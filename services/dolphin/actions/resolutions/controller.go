package resolutions

import (
	"github.com/mrNobody95/Gate/pkg/grpcext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/dolphin/actions/viewHelpers"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"net/http"
)

type resolutionController struct {
	resolutionService brokerageApi.ResolutionServiceClient
}

func newResolutionController() resolutionController {
	return resolutionController{
		resolutionService: brokerageApi.NewResolutionServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
	}
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
