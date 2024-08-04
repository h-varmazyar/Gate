package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/raven/configs"
	"github.com/h-varmazyar/Gate/services/raven/docs"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/chipmunk"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/core"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/eagle"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/network"
	"github.com/h-varmazyar/Gate/services/raven/internal/app/telegramBot"
	"github.com/h-varmazyar/gopack/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"net"
	"net/http"
	"time"
)

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server Gate server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	ctx := context.Background()
	logger := log.New()

	defaultConf := flag.Bool("default-configs", false, "run program with default config")
	flag.Parse()

	conf, err := loadConfigs(*defaultConf)
	if err != nil {
		log.Panicf("failed to read configs: %v", err)
	}

	//registerNgrok()

	initializeDocs(conf)

	logger.Infof("running %v(%v)", conf.ServiceName, conf.Version)

	initializeAndRegisterApps(ctx, logger, conf)
}

func loadConfigs(defaultConfig bool) (*Configs, error) {
	log.Infof("reding configs...")

	if defaultConfig {
		viper.SetConfigType("yaml")
		log.Infof("reading deafult configs")
		err := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		if err != nil {
			log.WithError(err).Error("read from default configs failed")
			return nil, err
		}
	} else {
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
					log.WithError(localErr).Error("read from default configs failed")
					return nil, localErr
				}
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

func initializeDocs(configs *Configs) {
	docs.SwaggerInfo.Title = "The Gate API document"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = configs.Version
	docs.SwaggerInfo.Host = configs.ApiExternalAddress
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func initializeAndRegisterApps(_ context.Context, logger *log.Logger, configs *Configs) {
	service.Serve(configs.HttpPort, func(lst net.Listener) error {
		router := mux.NewRouter(true)

		router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
		router.PathPrefix("/ping").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			pong := struct {
				Time string
			}{
				Time: time.Now().String(),
			}
			httpext.SendModel(res, req, http.StatusOK, pong)
		}).Methods(http.MethodGet)

		core.RegisterRoutes(router, logger, configs.CoreRouter)
		chipmunk.RegisterRoutes(router, logger, configs.ChipmunkRouter)
		eagle.RegisterRoutes(router, logger, configs.EagleRouter)
		telegramBot.RegisterRoutes(router, logger, configs.TelegramBotRouter)
		network.RegisterRoutes(router, logger, configs.NetworkRouter)

		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(configs.ServiceName, configs.Version)
}

func registerNgrok() {
	ctx := context.Background()
	l, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ngrok ingress url: ", l.Addr())
	http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from your ngrok-delivered Go app!")
	}))
}
