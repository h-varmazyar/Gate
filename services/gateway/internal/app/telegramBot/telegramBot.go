package telegramBot

import (
	gorilla "github.com/gorilla/mux"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	telegramBotApi "github.com/h-varmazyar/Gate/services/telegramBot/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	telegramBotService telegramBotApi.BotServiceClient
	logger             *log.Logger
}

func RegisterRoutes(router *gorilla.Router, logger *log.Logger, configs *Configs) {
	if controller == nil {
		telegramBotConn := grpcext.NewConnection(configs.TelegramBotAddress)
		controller = &Controller{
			telegramBotService: telegramBotApi.NewBotServiceClient(telegramBotConn),
			logger:             logger,
		}
	}
	telegramBotRouter := router.PathPrefix("/bot").Subrouter()

	telegramBotRouter.HandleFunc("/start", controller.start).Methods(http.MethodPost)
	telegramBotRouter.HandleFunc("/stop", controller.stop).Methods(http.MethodPost)
}

func (c *Controller) start(res http.ResponseWriter, req *http.Request) {
	log.Info("starting telegram bot...")
	if _, err := c.telegramBotService.Start(req.Context(), new(proto.Void)); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	log.Info("telegram bot started")
	httpext.SendCode(res, req, http.StatusOK)
}

func (c *Controller) stop(res http.ResponseWriter, req *http.Request) {
	log.Info("stoping telegram bot...")
	if _, err := c.telegramBotService.Stop(req.Context(), new(proto.Void)); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	log.Info("telegram bot stopped")
	httpext.SendCode(res, req, http.StatusOK)
}
