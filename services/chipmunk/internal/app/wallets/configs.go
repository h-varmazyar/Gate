package wallets

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/workers"
)

type Configs struct {
	ServiceConfigs *service.Configs `mapstructure:"service_configs"`
	BufferConfigs  *buffer.Configs  `mapstructure:"buffer_configs"`
	WorkerConfigs  *workers.Configs `mapstructure:"worker_configs"`
}
