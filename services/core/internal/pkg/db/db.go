package db

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/gormext"
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
			return nil, err
		}
		db.PostgresDB = postgres
	}

	return db, nil
}
