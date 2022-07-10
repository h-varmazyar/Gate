package configs

import (
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
	"time"
)

var Variables *Configs

type Configs struct {
	DBConnection              string        `env:"DATABASE_CONNECTION,required,file"`
	ServiceName               string        `env:"SERVICE_NAME,required"`
	Version                   string        `env:"VERSION,required"`
	IndicatorsWorkerHeartbeat time.Duration `env:"INDICATORS_WORKER_HEARTBEAT,required"`
	SignalsWorkerHeartbeat    time.Duration `env:"SIGNALS_WORKER_HEARTBEAT,required"`
	CandleBufferLength        int           `env:"CANDLE_BUFFER_LENGTH,required"`
	GrpcPort                  uint16        `env:"GRPC_PORT,required"`
	GrpcAddresses             struct {
		TelegramBot string `env:"TELEGRAM_BOT_GRPC_ADDRESS,required"`
		Chipmunk    string `env:"CHIPMUNK_GRPC_ADDRESS,required"`
		Brokerage   string `env:"BROKERAGE_GRPC_ADDRESS,required"`
	}
}

func LoadVariables() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
