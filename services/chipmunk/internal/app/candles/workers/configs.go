package workers

import "time"

type Configs struct {
	CoreAddress           string        `yaml:"core_address"`
	PrimaryDataQueue      string        `yaml:"primary_data_queue"`
	ConsumerCount         int           `yaml:"consumer_count"`
	MissedCandlesInterval time.Duration `yaml:"missed_candles_interval"`
}
