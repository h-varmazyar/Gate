package brokerages

import (
	"context"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
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
	MarketInfo(func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error)) (*repository.Market, error)
	//MarketList(func(ctx context.Context, request networkAPI.Request) networkAPI.Response)
}
