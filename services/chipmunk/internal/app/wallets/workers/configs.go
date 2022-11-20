package workers

import "time"

type Configs struct {
	CoreAddress          string        `yaml:"core_address"`
	WalletWorkerInterval time.Duration `yaml:"wallet_worker_interval"`
}
