package strategies

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"time"
)

type Strategy struct {
	HFT                   bool `yaml:"hft"`
	Markets               []models.Market
	StopLoss              float64
	MinBenefit            float64
	MaxBenefit            float64
	PrimaryAmount         float64
	CurrentAmount         float64
	PrimaryCurrency       models.Currency
	CurrentCurrency       models.Currency
	MaxDailyBenefit       float64
	ReservePercentage     float64
	BufferedCandleCount   int `yaml:"bufferedCandleCount"`
	IndicatorCalcLength   int
	IndicatorUpdatePeriod time.Duration `yaml:"indicatorUpdatePeriod"`
}

func (s *Strategy) Validate() error {
	if len(s.Markets) == 0 {
		return errors.New("trading symbols must be declared")
	}
	if s.StopLoss == 0 || s.StopLoss > 100 {
		return errors.New("stop loss percentage must be between 1 and 100")
	}
	//if s.Brokerage == nil {
	//	return errors.New("you must declared working brokerage")
	//}
	//if err := s.Brokerage.Validate(); err != nil {
	//	return err
	//}
	//if len(s.PivotResolution) == 0 {
	//	return errors.New("pivot time frame must be declared")
	//}
	if s.PrimaryCurrency == "" {
		return errors.New("primary currency must be declared")
	}
	if s.ReservePercentage < 0 || s.ReservePercentage > 100 {
		return errors.New("reserve percentage must be between 0 and 100")
	}
	return nil
}
