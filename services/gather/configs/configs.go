package configs

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
)

type AppEnv string

const (
	ProductionEnv AppEnv = "production"
	StageEnv      AppEnv = "stage"
	DevelopEnv    AppEnv = "develop"
	LocalEnv      AppEnv = "local"
)

type Config struct {
	GinMode  string
	AppEnv   AppEnv
	AppDebug bool
	Locale   string
	LogLevel log.Level
	Tz       string
	GRPC     GRPC
	HTTP     HTTP
	Database gormext.Configs

	MarketUpdateWorker WorkerMarketUpdate
	LastCandleWorker   WorkerLastCandle
	SahamyabArchive    WorkerSahamyabArchive
	SahamyabStream     WorkerSahamyabStream
	PostSentimentCheck WorkerPostSentimentCheck
	TickerWorker       WorkerTicker

	CandleBuffer    CandleBuffer
	CoreAdapter     CoreAdapter
	CoinexAdapter   CoinexAdapter
	SahamyabAdapter SahamyabAdapter
	Nats            Nats
}

type HTTP struct {
	APIHost string
	APIPort int
}

type GRPC struct {
	Port int
}
