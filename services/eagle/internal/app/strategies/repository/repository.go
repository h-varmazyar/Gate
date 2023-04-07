package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "strategies"

type StrategyRepository interface {
	Save(strategy *entity.Strategy) error
	Return(strategyID uuid.UUID) (*entity.Strategy, error)
	ReturnActives(ctx context.Context) ([]*entity.Strategy, error)
	ReturnIndicators(strategyID uuid.UUID) ([]*entity.StrategyIndicator, error)
	List() ([]*entity.Strategy, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (StrategyRepository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return NewStrategyPostgresRepository(ctx, logger, db.PostgresDB)
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
			if err := tx.AutoMigrate(new(entity.Strategy)); err != nil {
				return err
			}
			if err := tx.AutoMigrate(new(entity.StrategyIndicator)); err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				Table:       tableName,
				Tag:         "v1.0.0",
				Description: "create strategies and strategy_indicators tables",
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
