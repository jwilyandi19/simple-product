package main

import (
	"log"

	"github.com/jwilyandi19/simple-product/external/db"
	"github.com/jwilyandi19/simple-product/handler"
	"github.com/jwilyandi19/simple-product/helper"
	productRepo "github.com/jwilyandi19/simple-product/repository/product"
	productUsecase "github.com/jwilyandi19/simple-product/usecase/product"

	"github.com/jwilyandi19/simple-product/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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

	//redisConn := redis.InitRedisConnection(redisConfig)

	mainServer := echo.New()
	mainServer.Use(middleware.Recover())

	productRoutes := mainServer.Group("product")
	userRoutes := mainServer.Group("user")
	orderRoutes := mainServer.Group("order")

	productRepo := productRepo.NewProductRepository(dbConn)

	productUsecase := productUsecase.NewProductUsecase(productRepo)

	productHandler := handler.NewProductHandler(productUsecase)

	server.InitProductServer(productRoutes, productHandler)
	server.InitUserServer(userRoutes)
	server.InitOrderServer(orderRoutes)

	port := ":8080"
	mainServer.Start(port)

}
