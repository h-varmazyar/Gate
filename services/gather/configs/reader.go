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
		Database: gormext.Configs{
			DbType:      gormext.Type(loadString("DB_TYPE")),
			Port:        uint16(loadUint("DB_PORT")),
			Host:        loadString("DB_HOST"),
			Username:    loadString("DB_USERNAME"),
			Password:    loadString("DB_PASSWORD"),
			Name:        loadString("DB_NAME"),
			IsSSLEnable: loadBool("DB_IS_SSL_ENABLE"),
		},
		MarketUpdateWorker: WorkerMarketUpdate{
			RunningTime: loadString("WORKER_MARKET_UPDATE_RUNNING_TIME"),
			S3Bucket:    loadString("S3_BUCKET"),
			S3AccessKey: loadString("S3_ACCESS_KEY"),
			S3SecretKey: loadString("S3_SECRET_KEY"),
			S3Endpoint:  loadString("S3_ENDPOINT"),
			S3Region:    loadString("S3_REGION"),
		},
		LastCandleWorker: WorkerLastCandle{
			RunningInterval: loadDuration("LAST_CANDLE_WORKER_RUNNING_INTERVAL"),
		},
		CandleBuffer: CandleBuffer{
			CandleBufferLength: loadInt("CANDLE_BUFFER_LENGTH"),
		},
		CoreAdapter: CoreAdapter{
			GrpcAddress: loadString("CORE_GRPC_ADDRESS"),
		},
	}
	cfg.LogLevel, err = log.ParseLevel(loadString("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
