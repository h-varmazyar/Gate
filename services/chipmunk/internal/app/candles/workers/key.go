package workers

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
)

type Runner struct {
	Platform   api.Platform
	Market     *chipmunkApi.Market
	Resolution *chipmunkApi.Resolution
}
