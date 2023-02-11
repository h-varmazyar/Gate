package candles

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
)

type Configs struct {
	ServiceConfigs *service.Configs `mapstructure:"service_configs"`
	WorkerConfigs  *workers.Configs `mapstructure:"worker_configs"`
}
