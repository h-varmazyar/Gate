package workers

import "time"

type Configs struct {
	CoreAddress              string        `yaml:"core_address"`
	PrimaryDataQueue         string        `yaml:"primary_data_queue"`
	ConsumerCount            int           `yaml:"consumer_count"`
	LastCandlesInterval      time.Duration `yaml:"last_candles_interval"`
	MissedCandlesInterval    time.Duration `yaml:"missed_candles_interval"`
	RedundantRemoverInterval time.Duration `yaml:"redundant_remover_interval"`
}
