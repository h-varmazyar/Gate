package workers

import "time"

type Configs struct {
	CoreAddress               string        `yaml:"core_address"`
	PrimaryDataWorkerInterval time.Duration `yaml:"primary_data_worker_interval"`
	PrimaryDataQueue          string        `yaml:"primary_data_queue"`
}
