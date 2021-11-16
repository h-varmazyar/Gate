package gormext

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

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
