package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/brokerage/configs"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Brokerages *BrokerageRepository
)

func InitializingDB() {
	db, err := gormext.Open(gormext.PostgreSQL, configs.Variables.DatabaseConnection)
	if err != nil {
		log.WithError(err).Fatal("can not load repository configs")
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx,
			gormext.UUIDExtension,
		); err != nil {
			return err
		}

		if err = tx.AutoMigrate(
			new(Brokerage),
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("failed to migrate database")
	}

	Brokerages = &BrokerageRepository{db: db}
}
