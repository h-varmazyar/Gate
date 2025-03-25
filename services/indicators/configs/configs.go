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
	GinMode         string
	AppEnv          AppEnv
	AppDebug        bool
	Locale          string
	LogLevel        log.Level
	Tz              string
	GRPC            GRPC
	HTTP            HTTP
	NatsURL         string
	Database        gormext.Configs
	ChipmunkAdapter ChipmunkAdapter
}

type HTTP struct {
	APIHost string
	APIPort int
}

type GRPC struct {
	Port int
}
