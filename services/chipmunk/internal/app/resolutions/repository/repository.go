package repository

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const tableName = "resolutions"

type ResolutionRepository interface {
	Set(resolution *entity.Resolution) error
	Return(id uuid.UUID) (*entity.Resolution, error)
	ReturnByDuration(duration time.Duration, platform api.Platform) (*entity.Resolution, error)
	List(platform api.Platform) ([]*entity.Resolution, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (ResolutionRepository, error) {
	if err := migration(ctx, db); err != nil {
		return nil, err
	}
	return NewResolutionPostgresRepository(ctx, logger, db.PostgresDB)
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
			err = tx.AutoMigrate(new(entity.Resolution))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				Table:       tableName,
				Tag:         "v1.0.0",
				Description: "create resolutions table",
			})
		}

		if _, ok := migrations["v1.0.1"]; !ok {
			err = tx.Model(new(entity.Resolution)).CreateInBatches(defaultResolutions(), 10).Error
			if err != nil {
				return err
			}

			newMigrations = append(newMigrations, &db.Migration{
				Table:       tableName,
				Tag:         "v1.0.1",
				Description: "insert default resolutions",
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

func defaultResolutions() []*entity.Resolution {
	resolutions := make([]*entity.Resolution, 0)

	resolutions = append(resolutions, &entity.Resolution{
		Platform: api.Platform_Coinex,
		Duration: time.Minute * 15,
		Label:    "Coinex 15 minutes",
		Value:    "15min",
	})

	return resolutions
}
