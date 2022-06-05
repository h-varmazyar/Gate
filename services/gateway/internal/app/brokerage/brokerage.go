package brokerage

import (
	"github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/brokerage/brokerages"
)

func RegisterRoutes(router *mux.Router) {
	brokerageRouter := router.PathPrefix("/brokerage").Subrouter()
	brokerages.ControllerInstance().RegisterRoutes(brokerageRouter)
}
