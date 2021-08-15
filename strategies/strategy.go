package strategies

import (
	"errors"
	"github.com/mrNobody95/Gate/api"
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
	HelperCallbackChannel chan api.OHLCResponse
	CallbackChannel       chan api.OHLCResponse
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
		lastCandleId := uint(0)
		err := candle.LoadLast()
		if err != nil {
			if err.Error() == "record not found" {
				candle.Time = time.Now().Add(-time.Hour * 24 * 365).Unix()
			}
		} else {
			lastCandleId = candle.ID
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
			}
		} else {
			lastCandleId = candle.ID
		}
		err = (&scheduler.Job{
			Name: scheduler.RangeOHLC,
			Args: map[string]interface{}{
				"last_candle": lastCandleId,
				"resolution":  s.PivotResolution[symbol],
				"start_from":  candle.Time,
				"brokerage":   s.Brokerage,
				"callback":    &s.CallbackChannel,
				"symbol":      symbol,
			},
		}).Enqueue()
		if err != nil {
			panic(err)
		}
		err = (&scheduler.Job{
			Name: scheduler.RangeOHLC,
			Args: map[string]interface{}{
				"last_candle": lastCandleId,
				"resolution":  s.HelperResolution[symbol],
				"start_from":  helperCandle.Time,
				"brokerage":   s.Brokerage,
				"callback":    &s.HelperCallbackChannel,
				"symbol":      symbol,
			},
		}).Enqueue()
	}
}

func (s *Strategy) CollectPeriodicData() {
	for _, symbol := range s.Symbols {
		err := (&scheduler.Job{
			Name:   scheduler.SingleOHLC,
			Period: s.PivotResolution[symbol].Duration,
			Args: map[string]interface{}{
				"resolution": s.PivotResolution[symbol],
				"brokerage":  s.Brokerage,
				"callback":   &s.CallbackChannel,
				"symbol":     symbol,
			},
		}).EnqueueContinuously()
		if err != nil {
			panic(err)
		}
		err = (&scheduler.Job{
			Name:   scheduler.SingleOHLC,
			Period: s.HelperResolution[symbol].Duration,
			Args: map[string]interface{}{
				"resolution": s.HelperResolution[symbol],
				"brokerage":  s.Brokerage,
				"callback":   &s.HelperCallbackChannel,
				"symbol":     symbol,
			},
		}).EnqueueContinuously()
	}
}

func (s *Strategy) CheckForNewData() {
	go func() {
		for data := range s.CallbackChannel {
			if len(data.Candles) == 1 {
				var err error
				last := len(s.BufferedCandles[data.Symbol]) - 1
				if data.Candles[0].Time == s.BufferedCandles[data.Symbol][last].Time {
					s.BufferedCandles[data.Symbol][last] = data.Candles[0]
					s.updateIndicators(data.Symbol, false)
					err = s.BufferedCandles[data.Symbol][last].Update()
				} else {
					s.BufferedCandles[data.Symbol] = s.BufferedCandles[data.Symbol][1:]
					s.BufferedCandles[data.Symbol][last] = data.Candles[0]
					s.updateIndicators(data.Symbol, false)
					err = s.BufferedCandles[data.Symbol][last].Create()
				}
				if err != nil {
					log.Error(err)
				}
			} else if data.ContinueLast > 0 {
				candle := models.Candle{}
				candle.ID = data.ContinueLast
				err := candle.Load()
				if err != nil {
					log.Error(err)
					continue
				}
				s.BufferedCandles[data.Symbol] = make([]models.Candle, len(data.Candles)+1)
				s.BufferedCandles[data.Symbol][0] = candle
				for i := 1; i < len(data.Candles)+1; i++ {
					s.BufferedCandles[data.Symbol][i] = data.Candles[i-1]
					s.updateIndicators(data.Symbol, false)
					err := s.BufferedCandles[data.Symbol][i].Create()
					if err != nil {
						log.Error(err)
					}
				}
			} else {
				s.BufferedCandles[data.Symbol] = data.Candles
				s.calculateIndicators(data.Symbol, false)
			}
		}
	}()
	go func() {
		for data := range s.HelperCallbackChannel {
			if len(data.Candles) == 1 {
				var err error
				last := len(s.BufferedHelperCandles[data.Symbol]) - 1
				if data.Candles[0].Time == s.BufferedHelperCandles[data.Symbol][last].Time {
					s.BufferedHelperCandles[data.Symbol][last] = data.Candles[0]
					s.updateIndicators(data.Symbol, true)
					err = s.BufferedHelperCandles[data.Symbol][last].Update()
				} else {
					s.BufferedHelperCandles[data.Symbol] = s.BufferedHelperCandles[data.Symbol][1:]
					s.BufferedHelperCandles[data.Symbol][last] = data.Candles[0]
					s.updateIndicators(data.Symbol, true)
					err = s.BufferedHelperCandles[data.Symbol][last].Create()
				}
				if err != nil {
					log.Error(err)
				}
			} else if data.ContinueLast > 0 {
				candle := models.Candle{}
				candle.ID = data.ContinueLast
				err := candle.Load()
				if err != nil {
					log.Error(err)
					continue
				}
				s.BufferedHelperCandles[data.Symbol] = make([]models.Candle, len(data.Candles)+1)
				s.BufferedHelperCandles[data.Symbol][0] = candle
				for i := 1; i < len(data.Candles)+1; i++ {
					s.BufferedHelperCandles[data.Symbol][i] = data.Candles[i-1]
					s.updateIndicators(data.Symbol, true)
					err := s.BufferedHelperCandles[data.Symbol][i].Create()
					if err != nil {
						log.Error(err)
					}
				}
			} else {
				s.BufferedHelperCandles[data.Symbol] = data.Candles
				s.calculateIndicators(data.Symbol, true)
			}
		}
	}()
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

func (s *Strategy) updateIndicators(symbol brokerages.Symbol, isHelper bool) {
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
		if err := s.IndicatorConfig.UpdateBollingerBand(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.IndicatorConfig.UpdateMACD(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.IndicatorConfig.UpdatePSAR(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.IndicatorConfig.UpdateRSI(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.IndicatorConfig.UpdateStochastic(); err != nil {
			log.Error(err)
		}
	}(&wg)

	wg.Wait()
	s.CheckIndicators(symbol, isHelper)
}

func (s *Strategy) calculateIndicators(symbol brokerages.Symbol, isHelper bool) {
	if isHelper {
		s.IndicatorConfig.Candles = s.BufferedHelperCandles[symbol]
	} else {
		s.IndicatorConfig.Candles = s.BufferedCandles[symbol]
	}
	go s.IndicatorConfig.CalculateADX()
	go func() {
		if err := s.IndicatorConfig.CalculateBollingerBand(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		if err := s.IndicatorConfig.CalculateMACD(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		if err := s.IndicatorConfig.CalculatePSAR(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		if err := s.IndicatorConfig.CalculateRSI(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		if err := s.IndicatorConfig.CalculateStochastic(); err != nil {
			log.Error(err)
		}
	}()
}
