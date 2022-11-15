package candles

import "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
}
