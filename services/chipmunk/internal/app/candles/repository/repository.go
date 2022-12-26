package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type CandleRepository interface {
	Save(candle *entity.Candle) error
	HardDelete(candle *entity.Candle) error
	BulkHardDelete(candleIDs []uuid.UUID) error
	BulkInsert(candles []*entity.Candle) error
	ReturnLast(marketID, resolutionID uuid.UUID) (*entity.Candle, error)
	ReturnList(marketID, resolutionID uuid.UUID, limit, offset int) ([]*entity.Candle, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (CandleRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Candle)); err != nil {
		return nil, err
	}
	return NewCandlePostgresRepository(ctx, logger, db.PostgresDB)
}
