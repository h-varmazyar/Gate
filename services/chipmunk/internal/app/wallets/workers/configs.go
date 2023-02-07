package workers

import "time"

type Configs struct {
	CoreAddress          string        `mapstructure:"core_address"`
	WalletWorkerInterval time.Duration `mapstructure:"wallet_worker_interval"`
}
