package wallets

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/workers"
)

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
	BufferConfigs  *buffer.Configs  `yaml:"buffer_configs"`
	WorkerConfigs  *workers.Configs `yaml:"worker_configs"`
}
