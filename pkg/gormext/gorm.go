package gormext

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

type Type string

const (
	Mariadb    Type = "mariadb"
	PostgreSQL Type = "postgreSQL"
)

func Open(configs Configs) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch configs.DbType {
	case Mariadb:
		dialector = mysql.Open(configs.generateDSN())
	case PostgreSQL:
		dialector = postgres.Open(configs.generateDSN())
	default:
		return nil, fmt.Errorf("db type not supported(%s)", configs.DbType)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil && strings.Contains(err.Error(), "SQLSTATE 3D000") {
		log.Infof("creating database %v", configs.Name)
		rootDB, err := gorm.Open(postgres.Open(configs.rootDSN()), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.WithError(err).Error("failed to create root DB")
			return nil, err
		}

		if err = rootDB.Exec(fmt.Sprintf("CREATE DATABASE \"%v\";", configs.Name)).Error; err != nil {
			log.WithError(err).Error("failed to create database")
			return nil, err
		}

		db, err = gorm.Open(dialector, &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
	}
	if err != nil {
		log.WithError(err).Error("failed to initialize DB")
		return nil, err
	}
	pdb, err := db.DB()
	if err != nil {
		log.WithError(err).Error("failed to get sql DB")
		return nil, err
	}
	pdb.SetConnMaxLifetime(time.Minute)
	pdb.SetMaxIdleConns(10)
	pdb.SetMaxOpenConns(200)

	return db, nil
}
