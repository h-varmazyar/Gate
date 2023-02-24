package platforms

import "github.com/h-varmazyar/Gate/services/core/internal/app/platforms/service"

type Configs struct {
	ServiceConfigs *service.Configs `mapstructure:"service_configs"`
}
