package configs

import (
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
)

var Variables *Configs

type Configs struct {
	ServiceName        string `env:"SERVICE_NAME,required"`
	Version            string `env:"VERSION,required"`
	GrpcPort           uint16 `env:"GRPC_PORT,required"`
	HttpPort           uint16 `env:"HTTP_PORT,required"`
	MaxLogsPerPage     int64  `env:"MAX_LOGS_PER_PAGE,required"`
	DatabaseConnection string `env:"DATABASE_CONNECTION,required,file"`
	GrpcAddresses      struct {
		Chipmunk string `env:"CHIPMUNK_GRPC_ADDRESS,required"`
		Eagle    string `env:"EAGLE_GRPC_ADDRESS,required"`
		Network  string `env:"NETWORK_GRPC_ADDRESS,required"`
	}
}

func Load() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
