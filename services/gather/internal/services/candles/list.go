package candles

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"golang.org/x/net/context"
)

func (s Service) List(ctx context.Context, marketID, resolutionID uint) ([]models.Candle, error) {
	candles := make([]models.Candle, 0)

	end := false
	offset := 0
	for !end {
		tmp, err := s.candlesRepo.All(ctx, marketID, resolutionID, offset)
		if err != nil {
			return nil, err
		}
		if len(tmp) == 0 {
			end = true
		}
		offset += len(tmp)

		candles = append(candles, tmp...)
	}

	return candles, nil
}
