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
			Running:     loadBool("WORKER_MARKET_UPDATE_RUNNING"),
			NeedWarmup:  loadBool("WORKER_MARKET_NEED_WARMUP"),
			RunningTime: loadString("WORKER_MARKET_UPDATE_RUNNING_TIME"),
			S3Bucket:    loadString("S3_BUCKET"),
			S3AccessKey: loadString("S3_ACCESS_KEY"),
			S3SecretKey: loadString("S3_SECRET_KEY"),
			S3Endpoint:  loadString("S3_ENDPOINT"),
			S3Region:    loadString("S3_REGION"),
		},
		LastCandleWorker: WorkerLastCandle{
			Running:         loadBool("LAST_CANDLE_WORKER_RUNNING"),
			RunningInterval: loadDuration("LAST_CANDLE_WORKER_RUNNING_INTERVAL"),
		},
		SahamyabArchive: WorkerSahamyabArchive{
			Running: loadBool("SAHAMYAB_ARCHIVE_RUNNING"),
		},
		PostSentimentCheck: WorkerPostSentimentCheck{
			Running:           loadBool("WORKER_POST_SENTIMENT_CHECK_RUNNING"),
			SocksProxyAddress: loadString("SOCKS_PROXY_ADDRESS"),
			GeminiAPIKey:      loadString("GEMINI_API_KEY"),
		},
		SahamyabStream: WorkerSahamyabStream{
			Running: loadBool("SAHAMYAB_STREAM_RUNNING"),
		},
		TickerWorker: WorkerTicker{
			Running:         loadBool("TICK_WORKER_RUNNING"),
			RunningInterval: loadDuration("TICKER_WORKER_RUNNING_INTERVAL"),
		},

		CandleBuffer: CandleBuffer{
			CandleBufferLength: loadInt("CANDLE_BUFFER_LENGTH"),
		},
		CoreAdapter: CoreAdapter{
			GrpcAddress: loadString("CORE_GRPC_ADDRESS"),
		},
		CoinexAdapter: CoinexAdapter{
			APIBaseURL:        loadString("COINEX_ADAPTER_API_BASE_URL"),
			BaseURL:           loadString("COINEX_ADAPTER_BASE_URL"),
			SocksProxyAddress: loadString("SOCKS_PROXY_ADDRESS"),
		},
		SahamyabAdapter: SahamyabAdapter{
			GuestBaseURL: loadString("SAHAMYAB_GUEST_BASE_URL"),
		},
		Nats: Nats{URL: loadString("NATS_URL")},
	}
	cfg.LogLevel, err = log.ParseLevel(loadString("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
