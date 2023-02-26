package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/core"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/telegramBot"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	conf := loadConfigs()
	logger := log.New()

	initializeAndRegisterApps(ctx, logger, conf)
}

func loadConfigs() *Configs {
	configs := new(Configs)
	confBytes, err := ioutil.ReadFile("../configs/config.yaml")
	if err != nil {
		log.WithError(err).Fatal("can not load yaml file")
	}
	if err = yaml.Unmarshal(confBytes, configs); err != nil {
		log.WithError(err).Fatal("can not unmarshal yaml file")
	}
	return configs
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, configs *Configs) {
	service.Serve(configs.HttpPort, func(lst net.Listener) error {
		router := mux.NewRouter(true)

		core.RegisterRoutes(router, logger, configs.CoreRouter)
		chipmunk.RegisterRoutes(router, logger, configs.ChipmunkRouter)
		eagle.RegisterRoutes(router, logger, configs.EagleRouter)
		telegramBot.RegisterRoutes(router, logger, configs.TelegramBotRouter)

		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(configs.ServiceName, configs.Version)

}
