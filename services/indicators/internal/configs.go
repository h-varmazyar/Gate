package internal

import (
	"github.com/h-varmazyar/Gate/services/indicators/internal/service"
	"github.com/h-varmazyar/Gate/services/indicators/internal/workers"
)

type Configs struct {
	ChipmunkAddress string          `mapstructure:"chipmunk_address"`
	ServiceConfigs  service.Configs `mapstructure:"service_configs"`
	WorkersConfigs  workers.Configs `mapstructure:"workers_configs"`
}
