package configs

import (
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/spf13/viper"
)

func load() (*Configs, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	cfg := &Configs{
		HTTP: HTTP{
			APIHost: loadString("HTTP_HOST"),
			APIPort: loadInt("HTTP_PORT"),
		},
		NatsURL: loadString("NATS_URL"),
		DB: gormext.Configs{
			DbType:      gormext.Type(loadString("DB_TYPE")),
			Port:        uint16(loadUint("DB_PORT")),
			Host:        loadString("DB_HOST"),
			Username:    loadString("DB_USERNAME"),
			Password:    loadString("DB_PASSWORD"),
			Name:        loadString("DB_NAME"),
			IsSSLEnable: loadBool("DB_IS_SSL_ENABLE"),
		},
	}
	return cfg, nil
}
