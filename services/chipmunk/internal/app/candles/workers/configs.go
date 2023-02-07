package workers

import "time"

type Configs struct {
	CoreAddress              string        `mapstructure:"core_address"`
	PrimaryDataQueue         string        `mapstructure:"primary_data_queue"`
	ConsumerCount            int           `mapstructure:"consumer_count"`
	LastCandlesInterval      time.Duration `mapstructure:"last_candles_interval"`
	MissedCandlesInterval    time.Duration `mapstructure:"missed_candles_interval"`
	RedundantRemoverInterval time.Duration `mapstructure:"redundant_remover_interval"`
}
