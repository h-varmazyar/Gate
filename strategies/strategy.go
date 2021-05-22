package strategies

import (
	"errors"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
)

type Strategy struct {
	IsHFT                 bool                 `json:"is_hft" xml:"is_hft"`                                   //
	Symbols               []brokerages.Symbol  `json:"symbols" xml:"symbols"`                                 //
	StopLoss              float64              `json:"stop_loss" xml:"stop_loss"`                             //every trading stop loss percentage
	Brokerage             brokerages.Brokerage `json:"brokerage" xml:"brokerage"`                             //
	MinBenefit            float64              `json:"min_benefit" xml:"min_benefit"`                         //minimum benefit percentage in each trades
	MaxBenefit            float64              `json:"max_benefit" xml:"max_benefit"`                         //maximum benefit percentage in each trades
	PrimaryAmount         float64              `json:"primary_amount" xml:"primary_amount"`                   //primary strategy amount of primary currency
	CurrentAmount         float64              `json:"current_amount" xml:"current_amount"`                   //currency amount after each trades
	PrimaryCandles        []models.Candle      `json:"primary_candles" xml:"primary_candles"`                 //strategy primary candles
	PivotTimeFrame        string               `json:"pivot_time_frame" xml:"pivot_time_frame"`               //
	HelperTimeFrame       []string             `json:"helper_time_frame" xml:"helper_time_frame"`             //
	BufferedCandles       []models.Candle      `json:"buffered_candles" xml:"buffered_candles"`               //strategy buffered candles
	PrimaryCurrency       brokerages.Currency  `json:"primary_currency" xml:"primary_currency"`               //
	CurrentCurrency       brokerages.Currency  `json:"current_currency" xml:"current_currency"`               //
	MaxDailyBenefit       float64              `json:"max_daily_benefit" xml:"max_daily_benefit"`             //maximum daily benefit percentage
	ReservePercentage     float64              `json:"reserve_percentage" xml:"reserve_percentage"`           //reserve percentage value for golden positions...
	CandleBufferLength    int                  `json:"candle_buffer_length" xml:"candle_buffer_length"`       //maximum candles length
	HelperBufferedCandles [][]models.Candle    `json:"helper_buffered_candles" xml:"helper_buffered_candles"` //strategy buffered candles
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
	if s.PivotTimeFrame == "" {
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
