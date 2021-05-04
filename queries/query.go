package queries

import (
	"user-app/models"
	"user-app/database"
	"fmt"
)


func InsertItem(items models.Items) (bool,error) {
	db := database.GetDBPool()
	fmt.Println(&items)
	err :=db.Table("items").Create(&items)
	if err.Error != nil {
		fmt.Println(err.Error)
		return false,err.Error
	}

	return true,nil
}

func InsertOrder(order models.Orders) (bool,error) {
	db := database.GetDBPool()
	err :=db.Table("orders").Create(&order)
	if err.Error != nil {
		fmt.Println(err.Error)
		return false,err.Error
	}

	return true,nil
}

func UpdateItem(tabels string,orderid uint,order models.Items) (bool,error) {
	db := database.GetDBPool()
	fmt.Println(order)
	err := db.Table(tabels).Where("order_id = ?",orderid).Update(&order)
	if err.Error !=nil{
		return false,err.Error
	}
	return true,nil
}

func UpdateOrder(tabels string,orderid uint,order models.Orders) (bool,error) {
	db := database.GetDBPool()

	err := db.Table(tabels).Where("order_id = ?",orderid).Update(&order)
	if err.Error !=nil{
		return false,err.Error
	}
	return true,nil
}

func DeleteOrder(orderid string) (bool,error)  {
	db := database.GetDBPool()
	var orders models.Orders
	err := db.Where("order_id = ?",orderid).Delete(&orders)

	if err.Error !=nil{
		return false,err.Error
	}
	return true,nil
}

func DeleteItem(orderid string) (bool,error)  {
	db := database.GetDBPool()
	var items models.Items
	err := db.Where("order_id = ?",orderid).Delete(&items)

	if err.Error !=nil{
		return false,err.Error
	}
	return true,nil
}



func GetAllItems() (models.Items, error) {
	db := database.GetDBPool()
	var items models.Items
	err :=db.Debug().Table("items").Find(&items)

	if err != nil{
		return items,err.Error
	}
	
	return items,nil
}


func GetAllOrder() (models.Orders, error) {
	db := database.GetDBPool()
	var order models.Orders
	err :=db.Debug().Table("orders").Find(&order)

	if err != nil{
		return order,err.Error
	}

	return order,nil
}
