package indicators

import "github.com/mrNobody95/Gate/models"

type macdConfig struct {
	Candles      []models.Candle
	fastLength   int
	slowLength   int
	signalLength int
	source       Source
}

type Source string

const (
	SourceOpen  = "open"
	SourceClose = "close"
	SourceHigh  = "high"
	SourceLow   = "low"
	SourceHL2   = "hl2"
	SourceHLC3  = "hlc3"
	SourceOHLC4 = "ohlc4"
)

func NewMacdConfig(fastLength, slowLength, signalLength int, source Source) *macdConfig {
	return &macdConfig{
		fastLength:   fastLength,
		slowLength:   slowLength,
		signalLength: signalLength,
		source:       source,
	}
}

func CalculateMacd(candles []models.Candle, appendCandles bool) error {

}
