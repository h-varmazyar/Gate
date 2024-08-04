package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	"gorm.io/gorm"
)

const tableName = "indicators"

type IndicatorRepository interface {
	Create(indicator *entity.Indicator) error
	Return(indicatorID uuid.UUID) (*entity.Indicator, error)
	//List(ctx context.Context, indicatorType chipmunkApi.IndicatorType) ([]*entity.Indicator, error)
}

//func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (IndicatorRepository, error) {
//	if err := migration(ctx, db); err != nil {
//		return nil, err
//	}
//	return NewIndicatorPostgresRepository(ctx, logger, db.PostgresDB)
//}

func migration(_ context.Context, dbInstance *db.DB) error {
	var err error
	migrations := make(map[string]struct{})
	tags := make([]string, 0)
	err = dbInstance.PostgresDB.Table(db.MigrationTable).Where("table_name = ?", tableName).Select("tag").Find(&tags).Error
	if err != nil {
		return err
	}

	for _, tag := range tags {
		migrations[tag] = struct{}{}
	}

	newMigrations := make([]*db.Migration, 0)
	err = dbInstance.PostgresDB.Transaction(func(tx *gorm.DB) error {
		if _, ok := migrations["v1.0.0"]; !ok {
			err = tx.AutoMigrate(new(entity.Indicator))
			if err != nil {
				return err
			}
			newMigrations = append(newMigrations, &db.Migration{
				TableName:   tableName,
				Tag:         "v1.0.0",
				Description: "create indicators table",
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
