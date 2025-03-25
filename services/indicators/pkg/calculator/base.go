package calculator

import (
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"sync"
)

type base struct {
	lock       sync.Mutex
	id         uint
	Market     *chipmunkAPI.Market
	Resolution *chipmunkAPI.Resolution
}
