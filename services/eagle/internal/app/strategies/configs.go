package strategies

import (
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/service"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
)

type Configs struct {
	ServiceConfigs  *service.Configs           `mapstructure:"service_configs"`
	AutomatedWorker *automatedStrategy.Configs `mapstructure:"automated_worker"`
	ChipmunkAddress string                     `mapstructure:"chipmunk_address"`
}
