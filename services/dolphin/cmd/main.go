package main

import (
	"github.com/h-varmazyar/Gate/services/dolphin/actions"
	"github.com/sirupsen/logrus"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		logrus.Fatal(err)
	}
}
