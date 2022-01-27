package actions

import (
	"fmt"
	"github.com/mrNobody95/Gate/services/dolphin/actions/brokerages"
	"github.com/mrNobody95/Gate/services/dolphin/actions/markets"
	"github.com/mrNobody95/Gate/services/dolphin/actions/resolutions"
	"github.com/mrNobody95/Gate/services/dolphin/actions/wallets"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"log"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/sirupsen/logrus"
)

var application *app.App

func App() *app.App {
	if application == nil {
		application = app.New(app.Options{
			Env:         configs.Variables.Environment,
			SessionName: "_backoffice_session",
		})

		fmt.Println("initializing app")

		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			panic("create logger directory failed")
		}

		file, err := os.OpenFile("logs/backoffice.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		application.AccessLogger = &logrus.Logger{
			Out:       file,
			Formatter: &logrus.TextFormatter{},
			Level:     logrus.InfoLevel,
		}

		application.GET("/logout", func(c app.Context) error {
			s := c.Session()
			s.Values["username"] = nil
			s.Save(c.Request(), c.Response())
			return c.Redirect("/")
		})

		application.GET("/", func(c app.Context) error {
			return c.Render(200, "dashboard")
		})

		brokerages.RegisterRoutes(application)
		resolutions.RegisterRoutes(application)
		markets.RegisterRoutes(application)
		wallets.RegisterRoutes(application)

		application.ServeFiles("/", packr.New("../public", "../public"))
	}

	return application
}
