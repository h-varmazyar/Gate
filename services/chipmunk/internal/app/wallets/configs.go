package wallets

import "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/service"

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
}
