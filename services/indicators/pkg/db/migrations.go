package db

import (
	"context"
	"gorm.io/gorm"
)

type Migration struct {
	gorm.Model
	TableName   string
	Tag         string
	Description string
}

const MigrationTable = "migrations"

func createMigrateTable(_ context.Context, db *DB) error {
	err := db.AutoMigrate(new(Migration))
	if err != nil {
		return err
	}
	return nil
}
