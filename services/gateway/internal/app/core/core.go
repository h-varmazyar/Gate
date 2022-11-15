package core

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/core/brokerages"
	f "github.com/h-varmazyar/Gate/services/gateway/internal/app/core/functions"
)

func RegisterRoutes(router *gorilla.Router) {
	coreRouter := router.PathPrefix("/core").Subrouter()
	brokerages.HandlerInstance().RegisterRoutes(coreRouter)
	f.HandlerInstance().RegisterRoutes(coreRouter)
}
