package strategies

import "github.com/google/uuid"

type Strategy interface {
	CheckForSignals(marketID uuid.UUID, marketName string) float64
}
