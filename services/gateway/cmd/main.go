package main

import (
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/brokerage"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/telegramBot"
	"github.com/h-varmazyar/gopack/mux"
	"net"
	"net/http"
)

func main() {
	configs.Load()

	service.Serve(configs.Variables.HttpPort, func(lst net.Listener) error {
		router := mux.NewRouter(true)
		//router.Use(httpext.Authorization)

		brokerage.RegisterRoutes(router)
		chipmunk.RegisterRoutes(router)
		eagle.RegisterRoutes(router)
		telegramBot.RegisterRoutes(router)

		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
