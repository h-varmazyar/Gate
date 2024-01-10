package internal

import "github.com/h-varmazyar/Gate/services/indicators/internal/service"

type Configs struct {
	ServiceConfigs service.Configs `mapstructure:"service_configs"`
}
