package gormext

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Type string

const (
	Mariadb    Type = "mariadb"
	PostgreSQL Type = "postgreSQL"
)

func Open(Type Type, dsn string) (*gorm.DB, error) {
	switch Type {
	case Mariadb:
		return gorm.Open(mysql.Open(dsn))
	case PostgreSQL:
		return gorm.Open(postgres.Open(dsn))
	default:
		return nil, fmt.Errorf("db type not supported(%s)", Type)
	}
}
