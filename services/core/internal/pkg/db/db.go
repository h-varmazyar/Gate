package db

import (
	"context"
	"gorm.io/gorm"
)

type DB struct {
	PostgresDB *gorm.DB
}

func NewDatabase(ctx context.Context, configs *Configs) (*DB, error) {
	db := new(DB)
	if configs.PostgresDSN != "" {
		postgres, err := newPostgres(ctx, configs.PostgresDSN)
		if err != nil {
			return nil, err
		}
		db.PostgresDB = postgres
	}
	
	return db, nil
}
