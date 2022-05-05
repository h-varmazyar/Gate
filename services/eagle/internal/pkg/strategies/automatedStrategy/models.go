package automatedStrategy

import (
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"sync"
)

type AssetBalancePool struct {
	Lock            *sync.Mutex
	Market          *brokerageApi.Market
	Total           float64
	Available       float64
	Sold            float64
	Running         bool
	IsBaseOrderDone bool
	AveragePrice    float64
	MakerFeeRate    float64
	TakerFeeRate    float64
}
