package indicators

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
)

func cloneCandles(input []*repository.Candle) []*repository.Candle {
	var cloned []*repository.Candle
	for _, candle := range input {
		c := *candle
		cloned = append(cloned, &c)
	}
	return cloned
}
