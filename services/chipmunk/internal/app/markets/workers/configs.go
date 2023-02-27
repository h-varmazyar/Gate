package workers

import "time"

type Configs struct {
	CoreAddress                    string        `mapstructure:"core_address"`
	MarketStatisticsWorkerInterval time.Duration `mapstructure:"market_statistics_worker_interval"`
}
