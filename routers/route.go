package routers

import (
	"github.com/labstack/echo/v4"
	"user-app/controllers"
)

func OrderRute(e *echo.Echo)  {
	aGrup := e.Group("/order")
	aGrup.POST("/create",controllers.CereteOrder)
	aGrup.GET("/",controllers.GetOrder)
	aGrup.PUT("/update",controllers.UpdateOrders)
	aGrup.DELETE("/:order_id",controllers.Delete)
}
