package models

import (
	"github.com/mrNobody95/gorm"
	"github.com/mrNobody95/gorm/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var db *gorm.DB

func init() {
	var err error
	tmp := os.Getenv("dbPort")
	port, err := strconv.Atoi(tmp)
	if err != nil {
		log.Panicf("db port cast failed: %v", err)
	}
	err, db = (&gorm.DatabaseConfig{
		Type:                      gorm.MYSQL,
		Username:                  os.Getenv("dbUsername"),
		Password:                  os.Getenv("dbPassword"),
		CharSet:                   os.Getenv("dbCharset"),
		Name:                      os.Getenv("dbName"),
		Host:                      os.Getenv("dbHost"),
		Port:                      port,
		SSLMode:                   false,
		LogLevel:                  logger.Silent,
		ColorfulLog:               true,
		IgnoreRecordNotFoundError: false,
	}).Initialize(&Brokerage{}, &Market{}, &Resolution{}, &Candle{})

	if err != nil {
		log.Panicf("initializing db failed:%v", err)
	}
}
