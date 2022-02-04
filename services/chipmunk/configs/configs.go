package configs

import (
	"github.com/mrNobody95/Gate/pkg/envext"
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
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

var Variables *Configs

type Configs struct {
	ServiceName           string        `env:"SERVICE_NAME,required"`
	Version               string        `env:"VERSION,required"`
	StorageProvider       string        `env:"STORAGE_PROVIDER,required"`
	DatabaseConnection    string        `env:"DATABASE_CONNECTION,required,file"`
	OHLCWorkerHeartbeat   time.Duration `env:"OHLC_WORKER_HEARTBEAT,required"`
	WalletWorkerHeartbeat time.Duration `env:"WALLET_WORKER_HEARTBEAT,required"`
	CandleBufferLength    int           `env:"CANDLE_BUFFER_LENGTH,required"`
	GrpcPort              uint16        `env:"GRPC_PORT,required"`
	GrpcAddresses         struct {
		Brokerage string `env:"BROKERAGE_GRPC_ADDRESS,required"`
		Network   string `env:"NETWORK_GRPC_ADDRESS,required"`
	}
}

func init() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
