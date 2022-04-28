package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	log "github.com/sirupsen/logrus"
)

var (
	Candles    *CandleRepository
	Indicators *IndicatorRepository
)

const (
	MariaDB = "mariadb"
)

func init() {
	switch configs.Variables.StorageProvider {
	case MariaDB:
		db, err := gormext.Open(gormext.Mariadb, configs.Variables.DatabaseConnection)
		if err != nil {
			log.WithError(err).Fatal("can not open db connection")
		}
		if err := db.AutoMigrate(new(Candle)); err != nil {
			log.WithError(err).Fatal("migration failed for candles")
		}
		if err := db.AutoMigrate(new(Indicator)); err != nil {
			log.WithError(err).Fatal("migration failed for indicators")
		}
		Candles = &CandleRepository{db: db}
		Indicators = &IndicatorRepository{db: db}
	}
}
