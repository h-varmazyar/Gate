package resolutions

import "github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	resolutionController := newResolutionController()
	orders := application.Group("/resolutions")
	orders.GET("/{brokerage_name}/list", resolutionController.showBrokerageResolutions)
}
