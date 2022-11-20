package indicators

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
)

func cloneCandles(input []*entity.Candle) []*entity.Candle {
	var cloned []*entity.Candle
	for _, candle := range input {
		c := *candle
		cloned = append(cloned, &c)
	}
	return cloned
}
