package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Clients *clientRepository
)

func InitializingDB(configs *Configs) {
	db, err := gormext.Open(gormext.PostgreSQL, configs.DBConnection)
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
			new(Client),
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("failed to migrate database")
	}

	Clients = &clientRepository{db: db}
}
