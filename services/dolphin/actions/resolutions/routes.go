package resolutions

import "github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	resolutionController := newResolutionController()
	orders := application.Group("/resolutions")
	orders.POST("/add", resolutionController.add)
	orders.GET("/list", resolutionController.list)
	orders.GET("/{brokerage_name}/list", resolutionController.showBrokerageResolutions)
}
