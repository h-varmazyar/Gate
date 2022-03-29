package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
)

var (
	Assets      *AssetRepository
	Markets     *MarketRepository
	Brokerages  *BrokerageRepository
	Resolutions *ResolutionRepository
	Strategies  *StrategyRepository
	Indicators  *IndicatorRepository
)

func LoadRepositories(dsn string) {
	db, err := gormext.Open(gormext.Mariadb, dsn)
	if err != nil {
		log.WithError(err).Fatal("can not load repository configs")
	}
	if err := db.AutoMigrate(new(Asset)); err != nil {
		log.WithError(err).Fatal("migration failed for asset")
	}
	if err := db.AutoMigrate(new(Indicator)); err != nil {
		log.WithError(err).Fatal("migration failed for indicator")
	}
	if err := db.AutoMigrate(new(Strategy)); err != nil {
		log.WithError(err).Fatal("migration failed for strategy")
	}
	if err := db.AutoMigrate(new(Brokerage)); err != nil {
		log.WithError(err).Fatal("migration failed for brokerage")
	}
	if err := db.AutoMigrate(new(Market)); err != nil {
		log.WithError(err).Fatal("migration failed for market")
	}
	if err := db.AutoMigrate(new(Resolution)); err != nil {
		log.WithError(err).Fatal("migration failed for resolution")
	}
	Assets = &AssetRepository{db: db}
	Markets = &MarketRepository{db: db}
	Brokerages = &BrokerageRepository{db: db}
	Resolutions = &ResolutionRepository{db: db}
	Strategies = &StrategyRepository{db: db}
	Indicators = &IndicatorRepository{db: db}
}
