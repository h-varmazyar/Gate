package configs

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
)

type AppEnv string

const (
	ProductionEnv AppEnv = "production"
	StageEnv      AppEnv = "stage"
	DevelopEnv    AppEnv = "develop"
	LocalEnv      AppEnv = "local"
)

func New() (*Configs, error) {
	cfg, err := load()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type Configs struct {
	DB      gormext.Configs
	HTTP    HTTP
	NatsURL string
}
