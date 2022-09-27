package services

import (
	dto "blog-api-golang/DTO"
	"net/http"

	"github.com/jinzhu/gorm"
)

func CreateBlog(blog *Blog) (*Blog, *gorm.DB) {
	db.NewRecord(*blog)
	dbres := db.Create(blog)

	if dbres.Error == nil {

		return blog, dbres
	}
	return nil, dbres
}

func UpdateBlog(blog *dto.BlogDTO) (*dto.BlogDTO, *gorm.DB) {
	modelBlog := MatchBlogDTOtoModel(*blog, &Blog{})
	dbres := db.Save(modelBlog)

	if dbres.Error == nil {
		return blog, dbres
	}
	return nil, dbres
}

func GetAllBlogs(r *http.Request) ([]dto.BlogDTO, *gorm.DB) {
	var blogsDTO []dto.BlogDTO
	var blogsModel []Blog
	dbres := db.Model(&Blog{}).Scopes(Paginate(r)).Find(&blogsModel)
	if dbres.RowsAffected > 0 {

		for i, blog := range blogsModel {
			comments, _ := GetAllCommentsByBlog(blog)
			blogsDTO = append(blogsDTO, MatchBlogModelToDTO(blog, &dto.BlogDTO{}))

			blogsDTO[i].Comments = comments
		}
	}
	return blogsDTO, dbres
}

func GetBlogById(id int64) (*dto.BlogDTO, *gorm.DB) {
	var getBlog Blog
	var blogDTO dto.BlogDTO
	dbres := db.Where("id = ?", id).Find(&getBlog)
	if dbres.RowsAffected > 0 {
		blogDTO = MatchBlogModelToDTO(getBlog, &dto.BlogDTO{})
		comments, _ := GetAllCommentsByBlog(getBlog)

		blogDTO.Comments = comments

	}
	return &blogDTO, dbres
}

func GetBlogsByAuthor(r *http.Request, username string) ([]Blog, *gorm.DB) {
	var getBlogs []Blog
	author, exists := FindUserByUsername(username)
	if exists {

		dbres := db.Model(&Blog{}).Where("author_id = ?", author.ID).Scopes(Paginate(r)).Find(&getBlogs)
		return getBlogs, dbres
	}

	return nil, nil

}

func DeleteBlog(id int64) (Blog, *gorm.DB) {
	var blog Blog
	var getBlog *dto.BlogDTO
	getBlog, _ = GetBlogById(id)

	dbres := db.Where("id = ?", id).Delete(blog)
	blog = MatchBlogDTOtoModel(*getBlog, &blog)
	if dbres.Error == nil {
		comments, _ := GetAllCommentsByBlog(blog)
		for _, comment := range comments {
			_, _ = DeleteComment(int64(comment.ID))
		}

	}
	return blog, dbres
}

func GetBlogBySlug(slug string) (*Blog, *gorm.DB) {
	var getBlog Blog
	dbres := db.Where("slug = ?", slug).Find(&getBlog)
	return &getBlog, dbres
}

func GetBlogBySlugUnscoped(slug string) (*Blog, *gorm.DB) {
	var getBlog Blog
	dbres := db.Unscoped().Where("slug = ?", slug).Find(&getBlog)
	return &getBlog, dbres
}

func ExistsBlogBySlug(slug string) bool {

	_, dbres := GetBlogBySlugUnscoped(slug)
	if dbres.RowsAffected > 0 {
		return true
	}

	return false
}

func MatchBlogDTOtoModel(dto dto.BlogDTO, model *Blog) Blog {
	model.ID = dto.ID
	model.CreatedAt = dto.CreatedAt
	model.UpdatedAt = dto.UpdatedAt
	model.Author_ID = dto.Author_ID
	model.Title = dto.Title
	model.Slug = dto.Slug
	model.Body = dto.Body
	model.Comments = nil

	return *model
}

func MatchBlogModelToDTO(model Blog, dto *dto.BlogDTO) dto.BlogDTO {
	dto.ID = model.ID
	dto.CreatedAt = model.CreatedAt
	dto.UpdatedAt = model.UpdatedAt
	dto.Author_ID = model.Author_ID
	author, _ := FindUserById(int(dto.Author_ID))
	dto.Author_Name = author.Username
	dto.Title = model.Title
	dto.Slug = model.Slug
	dto.Body = model.Body
	return *dto
}
