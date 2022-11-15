package functions

import "github.com/h-varmazyar/Gate/services/core/internal/app/functions/service"

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
}
