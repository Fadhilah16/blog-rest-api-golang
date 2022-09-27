package services

import (
	dto "blog-api-golang/DTO"

	"github.com/jinzhu/gorm"
)

func CreateComment(comment *Comment) (*Comment, *gorm.DB) {
	db.NewRecord(*comment)
	dbres := db.Create(comment)

	if dbres.Error == nil {

		return comment, dbres
	}
	return nil, dbres
}

func CreateCommentOnBlog(blog *dto.BlogDTO, comment *Comment, author *User) (*Comment, *gorm.DB) {
	comment.Author_ID = int64(author.ID)
	comment.Blog_ID = int64(blog.ID)

	_, dbres := CreateComment(comment)
	if dbres.Error == nil {

		return comment, dbres
	}
	return nil, dbres
}

func UpdateComment(comment *Comment) (*Comment, *gorm.DB) {
	dbres := db.Save(comment)

	if dbres.Error == nil {
		return comment, dbres
	}
	return nil, dbres
}

func DeleteComment(id int64) (Comment, *gorm.DB) {
	var comment Comment
	dbres := db.Where("id = ?", id).Delete(comment)
	return comment, dbres
}

func DeleteCommentByBlogId(id int64) (Comment, *gorm.DB) {
	var comment Comment
	dbres := db.Where("blog_id = ?", id).Delete(comment)
	return comment, dbres
}

func GetAllCommentsByBlog(blog Blog) ([]dto.CommentDTO, *gorm.DB) {
	var getComments []Comment
	var getCommentsDTO []dto.CommentDTO
	dbres := db.Model(&Comment{}).Where("blog_id = ?", blog.ID).Find(&getComments)

	for _, cm := range getComments {
		commentDTO := dto.CommentDTO{}
		commentDTO = MatchCommentModelToDTO(cm, dto.CommentDTO{})
		getCommentsDTO = append(getCommentsDTO, commentDTO)

	}
	return getCommentsDTO, dbres
}

func GetCommentById(id int) (*Comment, *gorm.DB) {
	var getComment Comment

	dbres := db.Where("id = ?", id).Find(&getComment)
	return &getComment, dbres
}

func MatchCommentDTOtoModel(dto dto.CommentDTO, model Comment) Comment {
	model.ID = dto.ID
	model.CreatedAt = dto.CreatedAt
	model.UpdatedAt = dto.UpdatedAt
	model.Author_ID = dto.Author_ID
	model.Blog_ID = dto.Blog_ID
	model.Body = dto.Body
	return model
}

func MatchCommentModelToDTO(model Comment, dto dto.CommentDTO) dto.CommentDTO {
	dto.ID = model.ID
	dto.CreatedAt = model.CreatedAt
	dto.UpdatedAt = model.UpdatedAt
	dto.Author_ID = model.Author_ID
	author, _ := FindUserById(int(dto.Author_ID))
	dto.Author_Username = author.Username
	dto.Author_Name = author.Name
	dto.Blog_ID = model.Blog_ID
	var blog Blog
	db.Where("id = ?", dto.Blog_ID).Find(&blog)
	dto.Blog_Slug = blog.Slug
	dto.Body = model.Body

	return dto

}
