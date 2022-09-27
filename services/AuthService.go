package services

import (
	dto "blog-api-golang/DTO"

	"github.com/jinzhu/gorm"
)

func CreateUser(user *User) (*User, *gorm.DB) {
	_, exists := FindUserByUsername(user.Username)

	if exists == true {

		return nil, nil
	}
	db.NewRecord(*user)
	dbres := db.Create(user)
	return user, dbres

}

func UpdateUser(user *User) (*User, *gorm.DB) {
	dbres := db.Save(user)

	if dbres.Error == nil {

		return user, dbres
	}
	return nil, dbres

}

func FindUserByUsername(username string) (*User, bool) {
	var user User
	dbres := db.Model(&User{}).Where("username=?", username).First(&user)

	if dbres.RowsAffected > 0 {
		return &user, true
	}
	return nil, false
}

func FindUserById(id int) (*User, bool) {

	var user User
	dbres := db.Model(&User{}).Where("id=?", id).First(&user)
	if dbres.RowsAffected > 0 {

		return &user, true
	}

	return nil, false
}

func MatchUserProperties(userData dto.Register) *User {
	var user User
	user.Name = userData.Name
	user.Username = userData.Username
	user.Password = userData.Password

	return &user
}

func FindAuthorByBlogId(id int) User {
	var author *User
	var blog *dto.BlogDTO
	blog, _ = GetBlogById(int64(id))
	author, _ = FindUserById(int(blog.Author_ID))
	return *author
}

func IsBlogAuthor(username string, id int) bool {
	user, exists := FindUserByUsername(username)

	author := FindAuthorByBlogId(id)

	if exists {
		if author.Username == user.Username {
			return true
		}
	}
	return false
}

func FindAuthorByCommentId(id int) User {
	var author *User
	var comment *Comment
	comment, _ = GetCommentById(id)
	author, _ = FindUserById(int(comment.Author_ID))
	return *author
}

func IsCommentAuthor(username string, id int) bool {
	user, exists := FindUserByUsername(username)

	author := FindAuthorByCommentId(id)

	if exists {
		if author.Username == user.Username {
			return true
		}
	}
	return false
}
