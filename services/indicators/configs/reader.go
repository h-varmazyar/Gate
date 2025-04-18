package configs

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Read() (*Config, error) {
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

	cfg := &Config{
		GinMode:  loadString(gin.EnvGinMode),
		AppEnv:   AppEnv(loadString("APP_ENV")),
		AppDebug: loadBool("APP_DEBUG"),
		Locale:   loadString("LOCALE"),
		Tz:       loadString("TZ"),
		GRPC:     GRPC{Port: loadInt("GRPC_PORT")},
		HTTP: HTTP{
			APIHost: loadString("HTTP_HOST"),
			APIPort: loadInt("HTTP_PORT"),
		},
		NatsURL: loadString("NATS_URL"),
		Database: gormext.Configs{
			DbType:      gormext.Type(loadString("DB_TYPE")),
			Port:        uint16(loadUint("DB_PORT")),
			Host:        loadString("DB_HOST"),
			Username:    loadString("DB_USERNAME"),
			Password:    loadString("DB_PASSWORD"),
			Name:        loadString("DB_NAME"),
			IsSSLEnable: loadBool("DB_IS_SSL_ENABLE"),
		},
		ChipmunkAdapter: ChipmunkAdapter{BaseURL: loadString("CHIPMUNK_BASE_URL")},
	}
	cfg.LogLevel, err = log.ParseLevel(loadString("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
