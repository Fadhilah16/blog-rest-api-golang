package models

import (
	"blog-api-golang/config"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Author_ID int64  `json:"author_id"`
	Blog_ID   int64  `json:"blog_id"`
	Body      string `json:"body"`
}

func init() {
	config.Connect()
	config.GetDB().AutoMigrate(&Comment{})
}
