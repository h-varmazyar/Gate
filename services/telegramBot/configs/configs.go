package configs

import (
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
)

var Variables *Configs

type Configs struct {
	DBConnection  string `env:"DATABASE_CONNECTION,required,file"`
	ServiceName   string `env:"SERVICE_NAME,required"`
	Version       string `env:"VERSION,required"`
	DebugMode     bool   `env:"DEBUG_MODE"`
	BotToken      string `env:"BOT_TOKEN,required,file"`
	GrpcPort      uint16 `env:"GRPC_PORT,required"`
	GrpcAddresses struct{}
}

func LoadVariables() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
