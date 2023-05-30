package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "nikhil:nik123@tcp(localhost:3306)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println("opening error")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
