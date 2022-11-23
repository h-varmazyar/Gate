package automatedStrategy

import (
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"sync"
)

type AssetBalancePool struct {
	Lock            *sync.Mutex
	Market          *chipmunkApi.Market
	Total           float64
	Available       float64
	Sold            float64
	Running         bool
	IsBaseOrderDone bool
	AveragePrice    float64
	MakerFeeRate    float64
	TakerFeeRate    float64
}
