package strategies

import "github.com/mrNobody95/Gate/models"

func (s *Strategy) RsiSignal(previous, current models.Candle) bool {
	return previous.RSI.RSI < 30 && current.RSI.RSI >= 30
}

func (s *Strategy) BollingerBandSignal(previous, current models.Candle, makerFeeRate, takerFeeRate float64) bool {
	if previous.Low < previous.LowerBond {
		return false
	}
	price := current.Close * (1 + makerFeeRate/100) * (1 + s.MinGainPercent/100) * (1 + takerFeeRate/100)
	return price < current.UpperBond
}

// StochasticSignal condition: rsi must be 14 3 10
func (s *Strategy) StochasticSignal(previous, current models.Candle) bool {
	if current.IndexD > 20 || previous.IndexK > 20 {
		return false
	}
	return previous.IndexK < current.IndexK
}
