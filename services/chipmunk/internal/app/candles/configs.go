package candles

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
)

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
	BufferConfigs  *buffer.Configs  `yaml:"buffer_configs"`
	WorkerConfigs  *workers.Configs `yaml:"worker_configs"`
}
