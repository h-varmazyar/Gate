package signals

import (
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/signals/service"
)

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
}
