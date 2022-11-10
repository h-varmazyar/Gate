package requests

import (
	networkService "github.com/h-varmazyar/Gate/services/network/internal/app/requests/service"
)

type Configs struct {
	ServiceConfigs *networkService.Configs `yaml:"service_configs"`
}
