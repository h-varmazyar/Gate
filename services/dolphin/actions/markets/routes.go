package markets

import "github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	controller := newMarketController()
	markets := application.Group("/markets")
	markets.GET("/list", controller.list)
	markets.POST("/add", controller.add)
	markets.GET("/{market_id}/show", controller.show)
	markets.GET("/{brokerage_name}/list", controller.showBrokerageMarkets)
}
