package repository

import (
	"github.com/mrNobody95/Gate/pkg/gormext"
	"github.com/mrNobody95/Gate/services/chipmunk/configs"
	log "github.com/sirupsen/logrus"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 02.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

var (
	Candles *CandleRepository
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
		Candles = &CandleRepository{DB: db}
	}
}
