package gormext

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Type string

const (
	Mariadb    Type = "mariadb"
	PostgreSQL Type = "postgreSQL"
)

func Open(Type Type, dsn string) (*gorm.DB, error) {
	switch Type {
	case Mariadb:
		return gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	case PostgreSQL:
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			return nil, err
		}
		pdb, err := db.DB()
		if err != nil {
			return nil, err
		}
		pdb.SetConnMaxLifetime(time.Minute)
		pdb.SetMaxIdleConns(10)
		pdb.SetMaxOpenConns(200)
		return db, nil
	default:
		return nil, fmt.Errorf("db type not supported(%s)", Type)
	}
}
