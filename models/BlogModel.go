package models

import (
	"blog-api-golang/config"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Author_ID int64     `json:"author_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug" gorm:"unique"`
	Body      string    `json:"body"`
	Comments  []Comment `json:"comments" gorm:"Foreignkey:Blog_ID;association_foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func init() {
	config.Connect()
	config.GetDB().AutoMigrate(&Blog{})
}
