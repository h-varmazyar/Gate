package buffer

import (
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"sync"
)

type Market struct {
	candles    []*repository.Candle
	indicators map[string][]interface{}
}

type IndicatorResp struct {
	ID    string
	Value interface{}
}

type markets struct {
	lock         *sync.Mutex
	list         map[string]Market
	BufferLength int
}

var Markets *markets

func init() {
	Markets = &markets{
		lock:         new(sync.Mutex),
		list:         make(map[string]Market),
		BufferLength: 400,
	}
}

func (m *markets) Update(marketName string, candle *repository.Candle, indicatorList ...IndicatorResp) {
	market, ok := m.list[marketName]
	if !ok {
		market = Market{
			candles:    make([]*repository.Candle, m.BufferLength),
			indicators: make(map[string][]interface{}),
		}
	}
	if market.candles[m.BufferLength-1].Time.Equal(candle.Time) {
		market.candles[m.BufferLength-1] = candle
	} else {
		market.candles = append(market.candles[1:], candle)
	}
	for _, indicator := range indicatorList {
		if market.candles[m.BufferLength-1].Time.Equal(candle.Time) {
			market.indicators[indicator.ID][m.BufferLength-1] = indicator.Value
		} else {
			market.indicators[indicator.ID] = append(market.indicators[indicator.ID][1:], indicator.Value)
		}
	}

	m.lock.Lock()
	m.list[marketName] = market
	m.lock.Unlock()
}

func (m *markets) GetLastNCandles(marketName string, n int) []*repository.Candle {
	if market, ok := m.list[marketName]; !ok {
		return nil
	} else if market.candles == nil {
		return nil
	} else {
		return market.candles[len(m.list[marketName].candles)-n:]
	}
}

func (m *markets) GetLastNIndicatorValue(marketName, indicatorID string, n int) []interface{} {
	if market, ok := m.list[marketName]; !ok {
		return nil
	} else if market.indicators == nil {
		return nil
	} else if market.indicators[indicatorID] == nil {
		return nil
	} else {
		return market.indicators[indicatorID][len(m.list[marketName].candles)-n:]
	}
}
