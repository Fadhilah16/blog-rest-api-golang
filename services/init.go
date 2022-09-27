package services

import (
	"blog-api-golang/config"
	"blog-api-golang/models"

	"github.com/jinzhu/gorm"
)

type Blog models.Blog
type User models.User
type Comment models.Comment

var db *gorm.DB = config.GetDB()
