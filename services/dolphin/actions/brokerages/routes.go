package brokerages

import "github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	brokerageController := newBrokerageController()
	orders := application.Group("/brokerages")
	orders.GET("/list", brokerageController.list)
	orders.GET("/{brokerage_id}/show", brokerageController.show)
	orders.GET("/{brokerage_id}/overview", brokerageController.overview)
	orders.GET("/add", brokerageController.showAddPage)
	orders.POST("/add", brokerageController.add)
}
