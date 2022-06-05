package eagle

import (
	"github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle/signals"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle/strategies"
)

func RegisterRoutes(router *mux.Router) {
	eagleRouter := router.PathPrefix("/eagle").Subrouter()
	signals.ControllerInstance().RegisterRoutes(eagleRouter)
	strategies.ControllerInstance().RegisterRoutes(eagleRouter)
}
