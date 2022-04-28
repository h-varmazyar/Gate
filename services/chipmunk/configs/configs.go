package configs

import (
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
	"time"
)

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
		Eagle     string `env:"EAGLE_GRPC_ADDRESS,required"`
		Network   string `env:"NETWORK_GRPC_ADDRESS,required"`
	}
}

func init() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
