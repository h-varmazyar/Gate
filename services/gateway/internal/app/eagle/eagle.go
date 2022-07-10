package eagle

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle/signals"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle/strategies"
)

func RegisterRoutes(router *gorilla.Router) {
	eagleRouter := router.PathPrefix("/eagle").Subrouter()
	signals.ControllerInstance().RegisterRoutes(eagleRouter)
	strategies.ControllerInstance().RegisterRoutes(eagleRouter)
}
