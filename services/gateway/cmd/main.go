package main

import (
	"bytes"
	"context"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/core"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/telegramBot"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	logger := log.New()

	conf, err := loadConfigs()
	if err != nil {
		log.Panic("failed to read configs")
	}

	logger.Infof("conf:%v", conf)

	logger.Infof("running %v(%v)", conf.ServiceName, conf.Version)

	initializeAndRegisterApps(ctx, logger, conf)
}

func loadConfigs() (*Configs, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")  //path for docker compose configs
	viper.AddConfigPath("../configs") //path for local configs
	if err := viper.ReadInConfig(); err != nil {
		localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		if localErr != nil {
			return nil, localErr
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func initializeAndRegisterApps(_ context.Context, logger *log.Logger, configs *Configs) {
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
