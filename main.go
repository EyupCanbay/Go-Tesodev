package main

import (
	"os"
	"tesodev/configs"
	"tesodev/middleware"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("APP_ENV") == "development" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.LogMiddleware)

	configs.ConnectDB()

	e.Logger.Fatal(e.Start(":3000"))
}
