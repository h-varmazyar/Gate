package actions

import (
	"fmt"
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

		//application.Use(func(next app.Handler) app.Handler {
		//	return func(c app.Context) error {
		//
		//		s := c.Session()
		//		if s.Values["username"] != nil {
		//			return next(c)
		//		}
		//
		//		username := c.Request().Form.Get("username")
		//		password := c.Request().Form.Get("password")
		//		if info, ok := configs.Variables.Users[username]; ok {
		//			if hasher.CheckPasswordHash(password, info.Password) {
		//				s.Values["username"] = username
		//				s.Values["role"] = info.Role
		//				s.Save(c.Request(), c.Response())
		//				return c.Redirect("/")
		//			}
		//		}
		//		if c.Request().Method == "POST" {
		//			c.Set("Error", errors.New("invalid Username or Password"))
		//		}
		//		return c.Render(200, "login")
		//	}
		//})

		application.GET("/logout", func(c app.Context) error {
			s := c.Session()
			s.Values["username"] = nil
			s.Save(c.Request(), c.Response())
			return c.Redirect("/")
		})

		application.GET("/", func(c app.Context) error {
			fmt.Println("in dashboard")
			return c.Render(200, "dashboard")
		})

		//dashboard.RegisterRoutes(application)
		//finance.RegisterRoutes(application.Group("/finance"))
		//marketing.RegisterRoutes(application.Group("/marketing"))
		//user.RegisterRoutes(application.Group("/user"))

		application.ServeFiles("/", packr.New("../public", "../public"))
	}

	return application
}
