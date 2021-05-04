package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/labstack/gommon/log"
)

var (
	host     = "127.0.0.1"
	user     = "zamzam"
	password = "teuingah"
	port     = "1433"
)
var dbSQLServer *gorm.DB

func StartDB() (*gorm.DB, error) {
	dbmy, err := gorm.Open("mssql", "sqlserver://"+user+":"+password+"@"+host+":"+port+"?database=orders_by")
	dbmy.DB().SetMaxIdleConns(10)
	dbmy.LogMode(true)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	dbSQLServer = dbmy
	return dbSQLServer, nil
}

func GetDBPool() *gorm.DB {
	return dbSQLServer
}
