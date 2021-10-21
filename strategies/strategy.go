package strategies

import (
	"github.com/mrNobody95/Gate/models"
	"time"
)

type Strategy struct {
	HFT                   bool `yaml:"hft"`
	StopLoss              float64
	MinGainPercent        float64 `yaml:"minGainPercent"` //between 0-100
	LossPercent           float64 `yaml:"lossPercent"`    //between 0-100
	MaxGainPercent        float64 `yaml:"maxGainPercent"` //between 0-100
	PrimaryAmount         float64
	CurrentAmount         float64
	PrimaryCurrency       models.Currency
	CurrentCurrency       models.Currency
	MaxDailyBenefit       float64
	ReservePercentage     float64
	BufferedCandleCount   int           `yaml:"bufferedCandleCount"`
	IndicatorUpdatePeriod time.Duration `yaml:"indicatorUpdatePeriod"`
}

func (s *Strategy) Validate() error {
	return nil
}
