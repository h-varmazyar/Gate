package migrations

import (
	"github.com/h-varmazyar/p3o/internal/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(new(entities.User)); err != nil {
		return err
	}

	if err := db.AutoMigrate(new(entities.Link)); err != nil {
		return err
	}

	if err := db.AutoMigrate(new(entities.Visit)); err != nil {
		return err
	}

	return nil
}
