package indicators

import "github.com/mrNobody95/Gate/models"

func SimpleMovingAverage(candles []models.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Close
	}
	return sum / float64(len(candles))
}

func TypicalPriceMovingAverage(candles []models.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += (candle.Close + candle.High + candle.Low) / 3
	}
	return sum / float64(len(candles))
}
