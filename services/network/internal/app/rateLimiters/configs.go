package rateLimiters

import networkService "github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters/service"

type Configs struct {
	ServiceConfigs *networkService.Configs `yaml:"service_configs"`
}
