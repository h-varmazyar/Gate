package IPs

import networkService "github.com/h-varmazyar/Gate/services/network/internal/app/IPs/service"

type Configs struct {
	ServiceConfigs *networkService.Configs `yaml:"service_configs"`
}
