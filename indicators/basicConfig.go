package indicators

import (
	"github.com/mrNobody95/Gate/models"
)

type basicConfig struct {
	Candles []models.Candle
	Length  int
}

func cloneCandles(candles []models.Candle) []models.Candle {
	return append([]models.Candle{}, candles...)
}
