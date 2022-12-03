package workers

import (
	"context"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"time"
)

type MissedCandles struct {
	db          repository.CandleRepository
	configs     *Configs
	markets     []*chipmunkApi.Market
	resolutions []*chipmunkApi.Resolution
	ctx         context.Context
	Started     bool
}

func NewMissedCandles(_ context.Context, db repository.CandleRepository, configs *Configs, markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) *MissedCandles {
	return &MissedCandles{
		db:          db,
		configs:     configs,
		markets:     markets,
		resolutions: resolutions,
	}
}

func (w *MissedCandles) Start() {
	if !w.Started {
		go w.run()
		w.Started = true
	}
}

func (w *MissedCandles) run() {
	ticker := time.NewTicker(w.configs.MissedCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:

		}
	}
}

func (w *MissedCandles) prepareMarkets() error {
	for _, market := range w.markets {
		err := w.prepareResolutions(market)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *MissedCandles) prepareResolutions(market *chipmunkApi.Market) error {
	for _, resolution := range w.resolutions {
		err := w.checkForMissedCandles(market, resolution)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *MissedCandles) checkForMissedCandles(market *chipmunkApi.Market, resolution *chipmunkApi.Resolution) error {
	resolutionID, err := uuid.Parse(resolution.ID)
	if err != nil {
		return err
	}
	marketID, err := uuid.Parse(market.ID)
	if err != nil {
		return err
	}
	candles, err := w.db.ReturnList(marketID, resolutionID, 0, 1000000)
	if err != nil {
		return err
	}
	for i := 1; i < len(candles); i++ {
		if candles[i-1].Time.Add(time.Duration(resolution.Duration)).Before(candles[i].Time) {
			//todo: fetch missed candles asynchronously from server
		}
	}
	return nil
}
