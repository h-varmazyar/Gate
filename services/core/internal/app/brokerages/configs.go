package brokerages

import "github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
}
