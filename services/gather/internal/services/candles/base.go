package candles

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type candlesRepository interface {
	All(ctx context.Context, marketID, resolutionID uint, offset int) ([]models.Candle, error)
}
type Service struct {
	logger      *log.Logger
	candlesRepo candlesRepository
}

func NewService(
	logger *log.Logger,
	candlesRepo candlesRepository,
) Service {
	return Service{
		logger:      logger,
		candlesRepo: candlesRepo,
	}
}
