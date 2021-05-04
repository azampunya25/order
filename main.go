package main

import (
	"user-app/database"
	"github.com/labstack/echo/v4"
	"user-app/routers"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"github.com/swaggo/echo-swagger"
	"user-app/docs"
	"fmt"
)


// @Host localhost:8080
// @BasePath /orders
// @Version 1.0
// @Title "ORDER"
func main()  {
	db,err := database.StartDB()
	if err != nil {
		log.Fatal("database error : ",err.Error())
	}
	e := echo.New()
	routers.OrderRute(e)
	docs.SwaggerInfo.Title = "Order "
	docs.SwaggerInfo.Description = "This is a Order API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", "localhost", "8080")
	docs.SwaggerInfo.BasePath = "/order"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Start(":8080")

	defer db.Close()

}
