package storage

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"sync"
)

type CandlePool struct {
	candles      []models.Candle
	Capacity     int
	marketId     uint16
	resolutionId uint
	lock         *sync.Mutex
}

func NewPool(capacity int, marketId uint16, resolutionId uint) (*CandlePool, error) {
	if capacity == 0 {
		return nil, errors.New("zero pool capacity not allowed")
	}
	pool := new(CandlePool)
	pool.Capacity = capacity
	pool.marketId = marketId
	pool.resolutionId = resolutionId
	//pool.candles = make([]models.Candle, capacity)
	pool.lock = new(sync.Mutex)
	return pool, nil
}

func (pool *CandlePool) GetLastCandle() *models.Candle {
	if len(pool.candles) == 0 {
		return nil
	}
	last := pool.candles[len(pool.candles)-1]
	return &last
}

func (pool *CandlePool) GetLastNCandle(n int) []models.Candle {
	if len(pool.candles) == 0 || n > len(pool.candles) {
		return nil
	}
	last := pool.candles[len(pool.candles)-n:]
	return last
}

func (pool *CandlePool) UpdateLastCandle(candle models.Candle) error {
	lastIndex := len(pool.candles) - 1
	if candle.Market.ID != pool.candles[lastIndex].Market.ID ||
		candle.Resolution.Id != pool.candles[lastIndex].Resolution.Id {
		return errors.New("the candle is not belongs to this pool")
	}
	pool.lock.Lock()
	defer pool.lock.Unlock()
	if pool.candles[lastIndex].Time.Unix() == candle.Time.Unix() {
		pool.candles[lastIndex] = candle
	} else if pool.candles[lastIndex].Time.Unix() < candle.Time.Unix() {
		go func(candle models.Candle) {
			err := candle.CreateOrUpdate()
			log.WithError(err).Error("saving new candle failed")
		}(pool.candles[lastIndex])
		pool.candles = append(pool.candles, candle)
		if len(pool.candles) > pool.Capacity {
			pool.candles = pool.candles[1:]
		}
	}
	return nil
}

func (pool *CandlePool) ImportNewCandles(candles []models.Candle) error {
	for i := 1; i < len(candles); i++ {
		if candles[i-1].Time.Unix() > candles[i].Time.Unix() {
			return errors.New("candles are not ascending")
		}
	}
	pool.lock.Lock()
	defer pool.lock.Unlock()
	if len(candles) > pool.Capacity {
		pool.candles = append(pool.candles, candles[len(candles)-pool.Capacity:]...)
	} else {
		pool.candles = append(pool.candles, candles...)
	}
	return nil
}

func (pool *CandlePool) Size() int {
	return len(pool.candles)
}
