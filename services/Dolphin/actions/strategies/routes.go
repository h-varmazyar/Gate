package strategies

import "github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	controller := newStrategyController()
	orders := application.Group("/strategies")
	orders.POST("", controller.create)
	orders.GET("/list", controller.list)
	orders.GET("/{strategy_id}", controller.returnByID)
}
