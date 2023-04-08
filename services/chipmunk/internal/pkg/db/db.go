package db

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DB struct {
	PostgresDB *gorm.DB
}

func NewDatabase(ctx context.Context, configs gormext.Configs) (*DB, error) {
	db := new(DB)
	if configs.DbType == gormext.PostgreSQL {
		postgres, err := newPostgres(ctx, configs)
		if err != nil {
			log.WithError(err).Error("failed to create new postgres")
			return nil, err
		}
		db.PostgresDB = postgres

		err = createMigrateTable(ctx, db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func createMigrateTable(_ context.Context, db *DB) error {
	err := db.PostgresDB.AutoMigrate(new(Migration))
	if err != nil {
		return err
	}
	return nil
}
