package assets

import "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets/service"

type Configs struct {
	ServiceConfigs service.Configs `mapstructure:"service_configs"`
}
