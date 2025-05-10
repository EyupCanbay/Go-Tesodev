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

func main() {
	e := echo.New()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(os.Stdout)
	e.Use(middleware.LogMiddleware)

	dbClient := configs.ConnectDB()
	collection := configs.GetCollection(dbClient)

	repo := &repo.ProductRepository{Collection: collection}
	productService := &services.ProductService{Repo: repo}
	productHandler := &handlers.ProductHandler{Services: productService}

	e.POST("/product", middleware.LogMiddleware(productHandler.CreateProduct))
	e.GET("/product/:product_id", middleware.LogMiddleware(productHandler.GetAProduct))
	e.GET("/product", middleware.LogMiddleware(productHandler.GetAllProduct))
	e.PUT("/product/:product_id", middleware.LogMiddleware(productHandler.UpdateProduct))

	e.Logger.Fatal(e.Start(":3000"))
}
