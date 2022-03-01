package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type BollingerBands struct {
	basicConfig
	Deviation int
	Length    int
	Source    Source
}

func NewBollingerBands(length, deviation int, source Source, marketName string) *BollingerBands {
	return &BollingerBands{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
		},
		Deviation: deviation,
		Length:    length,
		Source:    source,
	}
}

func (conf *BollingerBands) GetID() string {
	return conf.id.String()
}

func (conf *BollingerBands) Calculate(candles []*repository.Candle, response interface{}) error {
	values := make([]*BollingerBandsResponse, len(candles))
	if err := conf.validateBollingerBand(len(candles)); err != nil {
		return err
	}
	cloned := cloneCandles(candles)
	smaConf := MovingAverage{
		Source: conf.Source,
		Length: conf.Length,
	}
	sma, err := smaConf.sma(cloned)
	if err != nil {
		return err
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		variance := float64(0)
		ma := sma[i]
		for j := 1 + i - conf.Length; j <= i; j++ {
			sum := float64(0)
			switch conf.Source {
			case SourceOpen:
				sum = candles[j].Open
			case SourceHigh:
				sum = candles[j].High
			case SourceLow:
				sum = candles[j].Low
			case SourceClose:
				sum = candles[j].Close
			case SourceOHLC4:
				sum = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
			case SourceHLC3:
				sum = (candles[j].Low + candles[j].High + candles[j].Close) / 3
			case SourceHL2:
				sum = (candles[j].Low + candles[j].High) / 2
			}
			variance += math.Pow(ma-sum, 2)
		}
		variance /= float64(conf.Length)
		values[i].MA = ma
		values[i].UpperBand = ma + float64(conf.Deviation)*math.Sqrt(variance)
		values[i].LowerBand = ma - float64(conf.Deviation)*math.Sqrt(variance)
	}
	response = interface{}(values)
	return nil
}

func (conf *BollingerBands) Update() interface{} {
	candles := buffer.Markets.GetLastNCandles(conf.MarketName, conf.Length)

	smaConf := MovingAverage{
		Source: conf.Source,
		Length: conf.Length,
	}
	sma, err := smaConf.sma(candles)
	if err != nil {
		return nil
	}
	variance := float64(0)
	ma := sma[len(candles)-1]
	for j := 0; j < len(candles); j++ {
		sum := float64(0)
		switch conf.Source {
		case SourceOpen:
			sum = candles[j].Open
		case SourceHigh:
			sum = candles[j].High
		case SourceLow:
			sum = candles[j].Low
		case SourceClose:
			sum = candles[j].Close
		case SourceOHLC4:
			sum = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
		case SourceHLC3:
			sum = (candles[j].Low + candles[j].High + candles[j].Close) / 3
		case SourceHL2:
			sum = (candles[j].Low + candles[j].High) / 2
		}
		variance += math.Pow(ma-sum, 2)
	}
	variance /= float64(conf.Length)
	return &BollingerBandsResponse{
		UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
		LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
		MA:        ma,
	}
}

func (conf *BollingerBands) validateBollingerBand(length int) error {
	if conf.Length != conf.Length {
		return errors.New("bollinger band length must be equal to moving average length")
	}
	if length < conf.Length {
		return errors.New("Length must be bigger than or equal to candle length")
	}
	if conf.Deviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}
