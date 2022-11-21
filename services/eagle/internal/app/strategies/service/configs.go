package service

import "github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"

type Configs struct {
	Automated *automatedStrategy.Configs `yaml:"automated"`
}
