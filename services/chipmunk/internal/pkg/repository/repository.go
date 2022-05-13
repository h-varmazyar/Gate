package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	log "github.com/sirupsen/logrus"
)

var (
	Assets      *AssetRepository
	Candles     *CandleRepository
	Indicators  *IndicatorRepository
	Resolutions *ResolutionRepository
	Markets     *MarketRepository
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

		if err := db.AutoMigrate(new(Asset)); err != nil {
			log.WithError(err).Fatal("migration failed for assets")
		}
		if err := db.AutoMigrate(new(Candle)); err != nil {
			log.WithError(err).Fatal("migration failed for candles")
		}
		if err := db.AutoMigrate(new(Indicator)); err != nil {
			log.WithError(err).Fatal("migration failed for indicators")
		}
		if err := db.AutoMigrate(new(Resolution)); err != nil {
			log.WithError(err).Fatal("migration failed for resolutions")
		}
		if err := db.AutoMigrate(new(Market)); err != nil {
			log.WithError(err).Fatal("migration failed for markets")
		}
		Assets = &AssetRepository{db: db}
		Candles = &CandleRepository{db: db}
		Indicators = &IndicatorRepository{db: db}
		Resolutions = &ResolutionRepository{db: db}
		Markets = &MarketRepository{db: db}
	}
}
