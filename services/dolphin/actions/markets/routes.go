package markets

import "github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	controller := newMarketController()
	markets := application.Group("/markets")
	markets.GET("/{brokerage_name}/list", controller.showBrokerageMarkets)
}
