package indicators

import "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators/service"

type Configs struct {
	ServiceConfigs service.Configs `mapstructure:"service_configs"`
}
