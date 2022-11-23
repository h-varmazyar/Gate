package workers

import "time"

type Configs struct {
	CoreAddress                    string        `yaml:"core_address"`
	MarketStatisticsWorkerInterval time.Duration `yaml:"market_statistics_worker_interval"`
}
