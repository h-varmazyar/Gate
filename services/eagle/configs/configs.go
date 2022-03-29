package configs

import (
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
	"time"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 01.12.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

var Variables *Configs

type Configs struct {
	ServiceName               string        `env:"SERVICE_NAME,required"`
	Version                   string        `env:"VERSION,required"`
	IndicatorsWorkerHeartbeat time.Duration `env:"INDICATORS_WORKER_HEARTBEAT,required"`
	SignalsWorkerHeartbeat    time.Duration `env:"SIGNALS_WORKER_HEARTBEAT,required"`
	CandleBufferLength        int           `env:"CANDLE_BUFFER_LENGTH,required"`
	GrpcAddresses             struct {
		Vault    uint16 `env:"VAULT_GRPC,required"`
		Chipmunk uint16 `env:"CHIPMUNK_GRPC,required"`
		Eagle    uint16 `env:"EAGLE_GRPC,required"`
	}
}

func init() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
