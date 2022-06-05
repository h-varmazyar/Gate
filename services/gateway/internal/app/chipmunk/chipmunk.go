package chipmunk

import (
	"github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk/assets"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk/indicators"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk/markets"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk/resolutions"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk/wallets"
)

func RegisterRoutes(router *mux.Router) {
	chipmunkRouter := router.PathPrefix("/chipmunk").Subrouter()
	assets.ControllerInstance().RegisterRoutes(chipmunkRouter)
	indicators.ControllerInstance().RegisterRoutes(chipmunkRouter)
	markets.ControllerInstance().RegisterRoutes(chipmunkRouter)
	resolutions.ControllerInstance().RegisterRoutes(chipmunkRouter)
	wallets.ControllerInstance().RegisterRoutes(chipmunkRouter)
}
