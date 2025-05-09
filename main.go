package main

import (
	"os"
	"tesodev/configs"
	"tesodev/handlers"
	"tesodev/middleware"
	"tesodev/repo"
	"tesodev/services"

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

	dbClient := configs.ConnectDB()
	collection := configs.GetCollection(dbClient)

	repo := &repo.ProductRepository{Collection: collection}
	productService := &services.ProductService{Repo: repo}
	userHandler := &handlers.ProductHandler{Services: productService}

	e.POST("/product", middleware.LogMiddleware(userHandler.CreateProduct))
	e.Logger.Fatal(e.Start(":3000"))
}
