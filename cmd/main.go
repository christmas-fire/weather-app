package main

import (
	"os"

	"github.com/christmas-fire/weather-app/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	app := app.NewApp()

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
