package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "candles"

type CandleRepository interface {
	Save(candle *entity.Candle) error
	HardDelete(candle *entity.Candle) error
	BulkHardDelete(candleIDs []uuid.UUID) error
	BulkInsert(candles []*entity.Candle) error
	ReturnLast(marketID, resolutionID uuid.UUID) (*entity.Candle, error)
	ReturnList(marketID, resolutionID uuid.UUID, limit, offset int) ([]*entity.Candle, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (CandleRepository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return NewCandlePostgresRepository(ctx, logger, db.PostgresDB)
}

func migration(_ context.Context, dbInstance *db.DB) error {
	var err error
	migrations := make(map[string]interface{})
	err = dbInstance.PostgresDB.Table(db.MigrationTable).Where("table = ?", tableName).Select("tag").Find(&migrations).Error
	if err != nil {
		return err
	}
	newMigrations := make([]*db.Migration, 0)
	err = dbInstance.PostgresDB.Transaction(func(tx *gorm.DB) error {
		if _, ok := migrations["v1.0.0"]; !ok {
			err = tx.AutoMigrate(new(entity.Candle))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				Table:       tableName,
				Tag:         "v1.0.0",
				Description: "create candles table",
			})
		}

		err = tx.Model(new(db.Migration)).CreateInBatches(newMigrations, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
