package strategies

import (
	"errors"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/scheduler"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Strategy struct {
	IsHFT                 bool                                    `json:"is_hft" xml:"is_hft"`                 //
	Symbols               []brokerages.Symbol                     `json:"symbols" xml:"symbols"`               //
	StopLoss              float64                                 `json:"stop_loss" xml:"stop_loss"`           //every trading stop loss percentage
	MinBenefit            float64                                 `json:"min_benefit" xml:"min_benefit"`       //minimum benefit percentage in each trades
	MaxBenefit            float64                                 `json:"max_benefit" xml:"max_benefit"`       //maximum benefit percentage in each trades
	PrimaryAmount         float64                                 `json:"primary_amount" xml:"primary_amount"` //primary strategy amount of primary currency
	CurrentAmount         float64                                 `json:"current_amount" xml:"current_amount"` //currency amount after each trades
	Brokerage             brokerages.Brokerage                    `json:"brokerage" xml:"brokerage"`
	PivotResolution       map[brokerages.Symbol]models.Resolution `json:"pivot_resolution" xml:"pivot_resolution"`               //
	BufferedCandles       map[brokerages.Symbol][]models.Candle   `json:"buffered_candles" xml:"buffered_candles"`               //strategy buffered candles
	PrimaryCurrency       brokerages.Currency                     `json:"primary_currency" xml:"primary_currency"`               //
	CurrentCurrency       brokerages.Currency                     `json:"current_currency" xml:"current_currency"`               //
	MaxDailyBenefit       float64                                 `json:"max_daily_benefit" xml:"max_daily_benefit"`             //maximum daily benefit percentage
	HelperResolution      map[brokerages.Symbol]models.Resolution `json:"helper_resolution" xml:"helper_resolution"`             //
	ReservePercentage     float64                                 `json:"reserve_percentage" xml:"reserve_percentage"`           //reserve percentage value for golden positions...
	CandleBufferLength    int                                     `json:"candle_buffer_length" xml:"candle_buffer_length"`       //maximum candles length
	BufferedHelperCandles map[brokerages.Symbol][]models.Candle   `json:"buffered_helper_candles" xml:"buffered_helper_candles"` //strategy buffered candles
	IndicatorConfig       indicators.IndicatorConfig
	DataChannel           chan struct {
		Symbol   brokerages.Symbol
		IsHelper bool
	}
}

func (s *Strategy) Validate() error {
	if len(s.Symbols) == 0 {
		return errors.New("trading symbols must be declared")
	}
	if s.StopLoss == 0 || s.StopLoss > 100 {
		return errors.New("stop loss percentage must be between 1 and 100")
	}
	if s.Brokerage == nil {
		return errors.New("you must declared working brokerage")
	}
	if err := s.Brokerage.Validate(); err != nil {
		return err
	}
	if len(s.PivotResolution) == 0 {
		return errors.New("pivot time frame must be declared")
	}
	if s.PrimaryCurrency == "" {
		return errors.New("primary currency must be declared")
	}
	if s.ReservePercentage < 0 || s.ReservePercentage > 100 {
		return errors.New("reserve percentage must be between 0 and 100")
	}
	return nil
}

func (s *Strategy) CollectPrimaryData() {
	for _, symbol := range s.Symbols {
		candle := models.Candle{
			Symbol:     symbol,
			Resolution: s.PivotResolution[symbol],
			Brokerage:  s.Brokerage.GetName(),
		}
		err := candle.LoadLast()
		if err != nil {
			if err.Error() == "record not found" {
				candle.Time = time.Now().Add(-time.Hour * 24 * 365).Unix()
			} else {
				continue
			}
		}
		helperCandle := models.Candle{
			Symbol:     symbol,
			Resolution: s.HelperResolution[symbol],
			Brokerage:  s.Brokerage.GetName(),
		}
		err = candle.LoadLast()
		if err != nil {
			if err.Error() == "record not found" {
				candle.Time = time.Now().Add(-time.Hour * 24 * 365).Unix()
			} else {
				continue
			}
		}
		err = (&scheduler.Job{
			Name: scheduler.RangeOHLC,
			Args: map[string]interface{}{"symbol": symbol,
				"resolution": s.PivotResolution[symbol],
				"start_from": candle.Time,
				"brokerage":  s.Brokerage,
				"callback":   &s.DataChannel,
			},
		}).Enqueue()
		if err != nil {
			panic(err)
		}
		err = (&scheduler.Job{
			Name: scheduler.RangeOHLC,
			Args: map[string]interface{}{"symbol": symbol,
				"resolution": s.HelperResolution[symbol],
				"start_from": helperCandle.Time,
				"brokerage":  s.Brokerage,
				"callback":   &s.DataChannel,
				"helper":     true,
			},
		}).Enqueue()
	}
}

func (s *Strategy) CollectPeriodicData() {
	for _, symbol := range s.Symbols {
		err := (&scheduler.Job{
			Name:   scheduler.SingleOHLC,
			Period: s.PivotResolution[symbol].Duration,
			Args: map[string]interface{}{"symbol": symbol,
				"resolution": s.PivotResolution[symbol],
				"brokerage":  s.Brokerage,
				"callback":   &s.DataChannel,
			},
		}).EnqueuePeriodically()
		if err != nil {
			panic(err)
		}
		err = (&scheduler.Job{
			Name:   scheduler.SingleOHLC,
			Period: s.HelperResolution[symbol].Duration,
			Args: map[string]interface{}{"symbol": symbol,
				"resolution": s.HelperResolution[symbol],
				"brokerage":  s.Brokerage,
				"callback":   &s.DataChannel,
				"helper":     true,
			},
		}).EnqueuePeriodically()
	}
}

func (s *Strategy) CheckForNewData() {
	go func() {
		for data := range s.DataChannel {
			s.calculateIndicators(data.Symbol, data.IsHelper)
		}
	}()
}

func (s *Strategy) calculateIndicators(symbol brokerages.Symbol, isHelper bool) {
	if isHelper {
		s.IndicatorConfig.Candles = s.BufferedHelperCandles[symbol]
	} else {
		s.IndicatorConfig.Candles = s.BufferedCandles[symbol]
	}
	var wg sync.WaitGroup
	wg.Add(6)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		s.IndicatorConfig.UpdateADX()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := s.IndicatorConfig.UpdateBollingerBand()
		if err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := s.IndicatorConfig.UpdateMACD()
		if err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := s.IndicatorConfig.UpdatePSAR()
		if err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := s.IndicatorConfig.UpdateRSI()
		if err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := s.IndicatorConfig.UpdateStochastic()
		if err != nil {
			log.Error(err)
		}
	}(&wg)

	wg.Wait()
	s.CheckIndicators(symbol, isHelper)
}

func (s *Strategy) CheckIndicators(symbol brokerages.Symbol, isHelper bool) {
	var last models.Candle
	if isHelper {
		last = s.BufferedHelperCandles[symbol][len(s.BufferedHelperCandles[symbol])-1]
	} else {
		last = s.BufferedCandles[symbol][len(s.BufferedCandles[symbol])-1]
	}
	if last.RSI.RSI < 30 {

	} else if last.RSI.RSI > 70 {

	}
}
