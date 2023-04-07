package repository

import (
	"context"
	"github.com/google/uuid"
	networkApi "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "rate_limiters"

type RateLimiterRepository interface {
	Create(brokerage *entity.RateLimiter) error
	Return(id uuid.UUID) (*entity.RateLimiter, error)
	List(Type networkApi.RateLimiterType) ([]*entity.RateLimiter, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (RateLimiterRepository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return NewRateLimiterPostgresRepository(ctx, logger, db.PostgresDB)
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
			err = tx.AutoMigrate(new(entity.RateLimiter))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				Table:       tableName,
				Tag:         "v1.0.0",
				Description: "create rate_limiters table",
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
