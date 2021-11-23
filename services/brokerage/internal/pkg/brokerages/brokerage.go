package brokerages

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"google.golang.org/grpc/codes"
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
* Date: 19.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Brokerage interface {
	WalletList(func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error)) ([]*repository.Wallet, error)
}

var (
	brokerages map[string]Brokerage
)

func init() {
	brokerages = make(map[string]Brokerage)
}

func Fetch(ctx context.Context, brokerage *brokerageApi.Brokerage) (Brokerage, error) {
	if brokerage.Status != api.StatusType_Enable {
		return nil, errors.NewWithSlug(ctx, codes.Unavailable, "brokerage is not enable")
	}
	if br, ok := brokerages[brokerage.ID]; ok {
		return br, nil
	}
	switch brokerage.Name {
	case brokerageApi.Names_Coinex:
		coinex := &Coinex{Auth: brokerage.Auth}
		brokerages[brokerage.ID] = coinex
		return coinex, nil
	default:
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "brokerage name is invalid")
	}
}
