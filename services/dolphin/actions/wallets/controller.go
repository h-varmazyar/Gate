package wallets

import (
	"github.com/mrNobody95/Gate/pkg/grpcext"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/dolphin/actions/viewHelpers"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
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
