package models

import (
	"blog-api-golang/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Blogs    []Blog `json:"blogs" gorm:"Foreignkey:Author_ID;association_foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func init() {
	config.Connect()
	config.GetDB().AutoMigrate(&User{})
}
