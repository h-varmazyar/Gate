package gormext

import "fmt"

type Configs struct {
	DbType      Type   `mapstructure:"type"`
	Port        uint16 `mapstructure:"port"`
	Host        string `mapstructure:"host"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	IsSSLEnable bool   `mapstructure:"is_ssl_enable"`
}

func (c *Configs) generateDSN() string {
	dsn := ""
	switch c.DbType {
	case PostgreSQL:
		dsn = fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", c.Username, c.Password, c.Host, "%v", c.Name, "%v")
		if c.Port != 0 {
			dsn = fmt.Sprintf(dsn, c.Port)
		}
		if c.IsSSLEnable {
			dsn = fmt.Sprintf(dsn, "Enable")
		} else {
			dsn = fmt.Sprintf(dsn, "disable")
		}
	}
	return dsn
}

func (c *Configs) rootDSN() string {
	dsn := ""
	switch c.DbType {
	case PostgreSQL:
		dsn = fmt.Sprintf("postgresql://%v:%v@%v:%v?sslmode=%v", c.Username, c.Password, c.Host, "%v", "%v")
		if c.Port != 0 {
			dsn = fmt.Sprintf(dsn, c.Port)
		}
		if c.IsSSLEnable {
			dsn = fmt.Sprintf(dsn, "Enable")
		} else {
			dsn = fmt.Sprintf(dsn, "disable")
		}
	}
	return dsn
}
