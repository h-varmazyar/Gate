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
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/test"
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
		log.Panicf("failed to read configs: %v", err)
	}

	logger.Infof("conf is: %v", conf)
	logger.Infof("core is: %v", conf.CoreRouter)
	logger.Infof("eagle is: %v", conf.EagleRouter)
	logger.Infof("chipmunk is: %v", conf.ChipmunkRouter)

	logger.Infof("running %v(%v)", conf.ServiceName, conf.Version)

	initializeAndRegisterApps(ctx, logger, conf)
}

func loadConfigs() (*Configs, error) {
	log.Infof("reding configs...")

	//os.Setenv("SERVICE_NAME", "gateway")
	//os.Setenv("VERSION", "v1.20.30")
	//os.Setenv("HTTP_PORT", "8080")
	//os.Setenv("CHIPMUNK_ROUTER_CHIPMUNK_ADDRESS", ":11000")
	//os.Setenv("CORE_ROUTER_CORE_ADDRESS", "core.gate.svc:10100")
	//os.Setenv("EAGLE_ROUTER_EAGLE_ADDRESS", ":12000")
	//os.Setenv("TELEGRAM_BOT_ROUTER_TELEGRAM_BOT_ADDRESS", ":14000")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("failed to read from env: %v", err)
		viper.AddConfigPath("./configs")  //path for docker compose configs
		viper.AddConfigPath("../configs") //path for local configs
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		if err = viper.ReadInConfig(); err != nil {
			log.Warnf("failed to read from yaml: %v", err)
			localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
			if localErr != nil {
				log.Infof("read from default env")
				return nil, localErr
			}
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		log.Errorf("faeiled unmarshal")
		return nil, err
	}

	return conf, nil
}

func initializeAndRegisterApps(_ context.Context, logger *log.Logger, configs *Configs) {
	service.Serve(configs.HttpPort, func(lst net.Listener) error {
		router := mux.NewRouter(true)

		test.RegisterRoutes(router, logger)
		core.RegisterRoutes(router, logger, configs.CoreRouter)
		chipmunk.RegisterRoutes(router, logger, configs.ChipmunkRouter)
		eagle.RegisterRoutes(router, logger, configs.EagleRouter)
		telegramBot.RegisterRoutes(router, logger, configs.TelegramBotRouter)

		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(configs.ServiceName, configs.Version)

}
