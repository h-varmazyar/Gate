package resolutions

import "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"

type Configs struct {
	ServiceConfigs *service.Configs `yaml:"service_configs"`
}
