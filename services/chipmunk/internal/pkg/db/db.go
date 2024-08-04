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

		log.Infof(configs.Name, configs.Host, configs.Password)
		err = createMigrateTable(ctx, db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func createMigrateTable(_ context.Context, db *DB) error {
	if err := db.PostgresDB.AutoMigrate(new(Migration)); err != nil {
		return err
	}

	if err := db.postsMigrations(); err != nil {
		return err
	}

	return nil
}
