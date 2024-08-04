package posts

import "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/posts/workers"

type Configs struct {
	WorkersConfigs workers.Configs `mapstructure:"workers_configs"`
}
