package wallets

import (
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/Dolphin/actions/viewHelpers"
	"github.com/h-varmazyar/Gate/services/Dolphin/configs"
	"github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"
	brokerageApi "github.com/h-varmazyar/Gate/services/core/api"
	"net/http"
)

type walletController struct {
	walletService brokerageApi.WalletServiceClient
}

func newWalletController() walletController {
	return walletController{
		walletService: brokerageApi.NewWalletServiceClient(grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)),
	}
}

func (c *walletController) list(ctx app.Context) error {
	//response, err := c.walletService.List(ctx, &brokerageApi.WalletListRequest{BrokerageName: brokerageApi.Names_All.String()})
	//if err != nil {
	//	return ctx.Error(http.StatusBadRequest, err)
	//}
	//ctx.Set("wallets", response.Wallets)
	return ctx.Render(http.StatusOK, "wallets/list", viewHelpers.Sum, viewHelpers.ResolutionLabel, viewHelpers.TimeStampFormat)
}
