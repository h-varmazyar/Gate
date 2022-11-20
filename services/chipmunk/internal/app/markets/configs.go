package markets

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/workers"
)

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
	WorkerConfigs  *workers.Configs `yaml:"worker_configs"`
}
