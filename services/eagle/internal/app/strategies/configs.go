package strategies

import (
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/service"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
)

type Configs struct {
	ServiceConfigs  *service.Configs           `yaml:"service_configs"`
	AutomatedWorker *automatedStrategy.Configs `yaml:"automated_worker"`
	ChipmunkAddress string                     `yaml:"chipmunk_address"`
}
