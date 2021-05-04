package models

import "time"

type Items struct {
	ItemID			uint	`gorm:"column:item_id" json:"itemId,omitempty"`
	ItemCode 		string	`gorm:"column:item_code" json:"itemCode"`
	Description 	string 	`gorm:"column:description" json:"description"`
	Quantity 		uint	`gorm:"column:quantity" json:"quantity"`
	OrderId 		uint 	`gorm:"column:order_id" json:"orderId,omitempty"`
}

type Orders struct {
	OrderId 		uint 		`gorm:"column:order_id"`
	CustomerName	string		`gorm:"column:customer_name"`
	OrderAt			time.Time	`gorm:"column:order_at"`
}


type Request struct {
	OrderId 		uint 		`json:"orderId,omitemty"`
	CustomerName 	string 		`json:"customerName"`
	OrderAt      	string		`json:"orderAt"`
	Items			[]Items		`json:"items"`
}

type Response struct {
	Status 	uint	`json:"status"`	
	Message	string	`json:"message"`
} 

type ErrorResponse struct {
	Status 	int		`json:"status"`
	Message string 	`json:"message"`
	Error 	string	`json:"error"`
}


