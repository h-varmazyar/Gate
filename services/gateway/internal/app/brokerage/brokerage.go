package brokerage

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/brokerage/brokerages"
)

func RegisterRoutes(router *gorilla.Router) {
	brokerageRouter := router.PathPrefix("/brokerage").Subrouter()
	brokerages.HandlerInstance().RegisterRoutes(brokerageRouter)
}
