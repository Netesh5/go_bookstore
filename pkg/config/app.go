package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connect() {
	dsn := "netesh:test@123/bookstore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}
