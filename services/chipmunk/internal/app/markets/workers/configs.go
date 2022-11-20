package workers

import "time"

type Configs struct {
	CoreAddress                    string        `yaml:"core_address"`
	OHLCWorkerHeartbeat            time.Duration `yaml:"ohlc_worker_heartbeat"`
	MarketStatisticsWorkerInterval time.Duration `yaml:"market_statistics_worker_interval"`
}
