package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {

	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog_api_golang?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
