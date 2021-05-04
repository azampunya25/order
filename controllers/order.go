package controllers

import (
	"github.com/labstack/echo/v4"
	"user-app/models"
	"net/http"
	"user-app/queries"
	"time"
	"fmt"
)

var req models.Request
var itemin models.Items
var orderin models.Orders

// Create Order godoc
// @Summary Create order
// @Description creating a new order
// @ID create-order
// @Accept  json,x-www-form-urlencoded
// @Produce  json
// @Param requestbody body models.Request true "request body json"
// @Success 200 {object} models.Request
// @Failure 500 {object} models.ErrorResponse
// @Router /create [post]
func CereteOrder(ctx echo.Context)  error{

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}

	date,err := time.Parse(time.RFC3339,req.OrderAt)
	if err != nil{
		fmt.Println(err.Error())
	}
	orderin.OrderId = req.OrderId
	orderin.OrderAt = date
	orderin.CustomerName = req.CustomerName
	for _,v := range req.Items {
		itemin = v
	}
	itemin.OrderId = req.OrderId
	_ ,err = queries.InsertItem(itemin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}

	_ ,err = queries.InsertOrder(orderin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}

	return ctx.JSON(http.StatusOK,req)
}

// Get Order godoc
// @Summary Getting all Order
// @Description getting all Order
// @ID get-Order
// @Produce  json
// @Success 200 {array} models.Request
// @Failure 500 {object} models.ErrorResponse
// @Router / [get]
func GetOrder(ctx echo.Context)  error{
	item,err := queries.GetAllItems()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}

	orderin,err = queries.GetAllOrder()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}


	req.CustomerName = orderin.CustomerName

	req.OrderAt = orderin.OrderAt.Format(time.RFC3339)
	items := make([]models.Items,0)
	req.Items =append(items,item)
	
	return ctx.JSON(http.StatusOK,req)
}

// Update Order godoc
// @Summary Update order
// @Description Update order
// @ID update-order
// @Accept  json,x-www-form-urlencoded
// @Produce  json
// @Param requestbody body models.Request true "request body json"
// @Success 200 {object} models.Request
// @Failure 500 {object} models.ErrorResponse
// @Router /update [put]
func UpdateOrders(ctx echo.Context)  error{

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}
	date,err := time.Parse(time.RFC3339,req.OrderAt)
	if err != nil{
		fmt.Println(err.Error())
	}
	orderin.OrderAt = date
	orderin.CustomerName = req.CustomerName
	for _,v := range req.Items {
		itemin = v
	}
	_,err =queries.UpdateItem("items",req.OrderId,itemin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}
	_,err = queries.UpdateOrder("orders",req.OrderId,orderin)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}
	return ctx.JSON(http.StatusOK,req)
}

// Delete Order godoc
// @Summary Delete Order
// @Description Delete Order
// @ID Delete-Order
// @Produce  json
// @Param order_id path string true "Order ID"
// @Success 200 {array} models.Request
// @Failure 500 {object} models.ErrorResponse
// @Router /:order_id [delete]
func Delete(ctx echo.Context)  error{

	param := ctx.Param("order_id")

	_,err :=queries.DeleteItem(param)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}
	_,err =queries.DeleteOrder(param)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,models.ErrorResponse{Status:500,Message:"Internal Server Error",Error:err.Error()})
	}



	return ctx.JSON(http.StatusOK,models.Response{Status:200,Message:"Order Berhasil di Delete"})
}
