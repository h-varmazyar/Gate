package configs

import (
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
)

var Variables *Configs

type Configs struct {
	ServiceName   string `env:"SERVICE_NAME,required"`
	Version       string `env:"VERSION,required"`
	HttpPort      uint16 `env:"HTTP_PORT,required"`
	GrpcAddresses struct {
		Core        string `env:"CORE_GRPC_ADDRESS,required"`
		Chipmunk    string `env:"CHIPMUNK_GRPC_ADDRESS,required"`
		Eagle       string `env:"EAGLE_GRPC_ADDRESS,required"`
		TelegramBot string `env:"TELEGRAM_BOT_ADDRESS,required"`
	}
}

func Load() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
