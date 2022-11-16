package repository

//import (
//	"github.com/h-varmazyar/Gate/pkg/gormext"
//	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
//	log "github.com/sirupsen/logrus"
//	"gorm.io/gorm"
//)
//
//var (
//	Assets      *AssetRepository
//	Candles     *CandleRepository
//	Indicators  *IndicatorRepository
//	Resolutions *ResolutionRepository
//	Markets     *MarketRepository
//)
//
//func InitializingDB() {
//	db, err := gormext.Open(gormext.PostgreSQL, configs.Variables.DatabaseConnection)
//	if err != nil {
//		log.WithError(err).Fatal("can not open db connection")
//	}
//
//	if err = db.Transaction(func(tx *gorm.DB) error {
//		if err = gormext.EnableExtensions(tx,
//			gormext.UUIDExtension,
//		); err != nil {
//			return err
//		}
//
//		if err = tx.AutoMigrate(
//			new(Asset),
//			new(Candle),
//			new(Indicator),
//			new(Resolution),
//			new(Market),
//		); err != nil {
//			return err
//		}
//		return nil
//	}); err != nil {
//		log.WithError(err).Fatal("failed to migrate database")
//	}
//
//	Assets = &AssetRepository{db: db}
//	Candles = &CandleRepository{db: db}
//	Indicators = &IndicatorRepository{db: db}
//	Resolutions = &ResolutionRepository{db: db}
//	Markets = &MarketRepository{db: db}
//}
