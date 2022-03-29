package wallets

import "github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/app"

func RegisterRoutes(application *app.App) {
	resolutionController := newWalletController()
	orders := application.Group("/wallets")
	//orders.POST("/add", resolutionController.add)
	orders.GET("/list", resolutionController.list)
}
