package db

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func newPostgres(_ context.Context, configs gormext.Configs) (*gorm.DB, error) {
	db, err := gormext.Open(configs)
	if err != nil {
		log.WithError(err).Error("failed to open database")
		return nil, err
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx,
			gormext.UUIDExtension,
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Error("failed to add extensions to database")
		return nil, err
	}

	return db, nil
}
