package main

import (
	"io"
	"os"

	"github.com/jwilyandi19/simple-product/external/db"
	"github.com/jwilyandi19/simple-product/handler"
	"github.com/jwilyandi19/simple-product/helper"
	productRepo "github.com/jwilyandi19/simple-product/repository/product"
	productUsecase "github.com/jwilyandi19/simple-product/usecase/product"
	log "github.com/sirupsen/logrus"

	userRepo "github.com/jwilyandi19/simple-product/repository/user"
	userUsecase "github.com/jwilyandi19/simple-product/usecase/user"

	orderRepo "github.com/jwilyandi19/simple-product/repository/order"
	orderUsecase "github.com/jwilyandi19/simple-product/usecase/order"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fileLog, err := os.OpenFile("simple_product.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fileLog.Close()

	mw := io.MultiWriter(fileLog, os.Stdout)

	log.SetOutput(mw)
	log.SetLevel(log.InfoLevel)
	config, err := helper.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config")
	}

	dbConfig := helper.DBConfig{
		Host:     config.DBHost,
		Password: config.DBPassword,
		Username: config.DBUsername,
		DB:       config.DBName,
	}

	// redisConfig := helper.RedisConfig{
	// 	Server:   config.RedisHost,
	// 	Password: config.RedisPassword,
	// }

	dbConn, err := db.InitDBConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	//redisConn := cache.InitCacheConnection(redisConfig)

	mainServer := echo.New()
	mainServer.Use(middleware.Recover())
	//mainServer.Use(middleware.Logger())

	productRoutes := mainServer.Group("product")
	userRoutes := mainServer.Group("user")
	orderRoutes := mainServer.Group("order")

	productRepo := productRepo.NewProductRepository(dbConn)
	userRepo := userRepo.NewUserRepository(dbConn)
	orderRepo := orderRepo.NewOrderRepository(dbConn)

	productUsecase := productUsecase.NewProductUsecase(productRepo)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	orderUsecase := orderUsecase.NewOrderUsecase(orderRepo)

	handler.NewProductHandler(productRoutes, productUsecase)
	handler.NewUserHandler(userRoutes, userUsecase)
	handler.NewOrderHandler(orderRoutes, orderUsecase)

	port := ":8080"
	mainServer.Start(port)

}
