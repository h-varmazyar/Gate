package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type BollingerBands struct {
	Deviation int
	Length    int
	Source    Source
	Values    []*BollingerBandsResponse
}

func (conf *BollingerBands) Calculate(candles []*repository.Candle) error {
	conf.Values = make([]*BollingerBandsResponse, len(candles))
	if err := conf.validateBollingerBand(len(candles)); err != nil {
		return err
	}
	cloned := cloneCandles(candles)
	smaConf := MovingAverage{
		Source: conf.Source,
		Length: conf.Length,
		Values: make([]*MovingAverageResponse, len(cloned)),
	}
	err := smaConf.sma(cloned)
	if err != nil {
		return err
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		variance := float64(0)
		ma := smaConf.Values[i].Simple
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
		conf.Values[i].MA = ma
		conf.Values[i].UpperBand = ma + float64(conf.Deviation)*math.Sqrt(variance)
		conf.Values[i].LowerBand = ma - float64(conf.Deviation)*math.Sqrt(variance)
	}
	return nil
}

func (conf *BollingerBands) Update(candles []*repository.Candle) (*BollingerBandsResponse, error) {
	response := new(BollingerBandsResponse)
	cloned := cloneCandles(candles[len(candles)-conf.Length:])
	smaConf := MovingAverage{
		Source: conf.Source,
		Length: conf.Length,
		Values: make([]*MovingAverageResponse, len(cloned)),
	}
	err := smaConf.sma(cloned)
	if err != nil {
		return nil, err
	}
	variance := float64(0)
	ma := smaConf.Values[len(cloned)-1].Simple
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
	//conf.lock.Lock()
	//candles[len(candles)-1].BollingerBands.MovingAverage = ma
	//candles[len(candles)-1].BollingerBands.UpperBand = ma + float64(conf.Deviation)*math.Sqrt(variance)
	//candles[len(candles)-1].BollingerBands.LowerBand = ma - float64(conf.Deviation)*math.Sqrt(variance)
	//conf.lock.Unlock()
	response.MA = ma
	response.UpperBand = ma + float64(conf.Deviation)*math.Sqrt(variance)
	response.LowerBand = ma - float64(conf.Deviation)*math.Sqrt(variance)
	return response, nil
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
