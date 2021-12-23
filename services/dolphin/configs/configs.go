package configs

import (
	"github.com/mrNobody95/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
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
	ServiceName      string `env:"SERVICE_NAME,required"`
	Version          string `env:"VERSION,required"`
	Host             string `env:"ADDR,required"`
	Port             uint16 `env:"PORT,required"`
	TLS              bool   `env:"TLS,required"`
	Environment      string `env:"GO_ENV,required"`
	ApiPrefix        string `env:"API_PREFIX,required"`
	AssetsAgeVarName int64  `env:"ASSETS_MAX_AGE,required"`
	LogLevel         string `env:"LOG_LEVEL"`
	GrpcAddresses    struct {
		Brokerage string `env:"BROKERAGE_GRPC,required"`
		Vault     string `env:"VAULT_GRPC,required"`
		Chipmunk  string `env:"CHIPMUNK_GRPC,required"`
		Eagle     string `env:"EAGLE_GRPC,required"`
	}
}

func init() {
	Variables = new(Configs)
	if err := envext.Load(Variables); err != nil {
		log.WithError(err).Panic("load env failed")
	}
}
