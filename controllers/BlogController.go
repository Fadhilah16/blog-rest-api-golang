package controllers

import (
	dto "blog-api-golang/DTO"
	"blog-api-golang/models"
	"blog-api-golang/services"
	"blog-api-golang/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {

	blog := &services.Blog{}
	utils.ParseBody(r, blog)
	if blog.Slug == "" {
		blog.Slug = utils.GenerateUniqueSlug(blog.Title)
	}
	author, _ := services.FindUserByUsername(r.Header.Get("username"))
	blog.Author_ID = int64(author.ID)

	_, dbres := services.CreateBlog(blog)
	var response dto.Response
	if dbres.Error == nil {

		response.Status = http.StatusOK
		response.Message = append(response.Message, "Blog successfully created")
		response.Entity = blog
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to create blog")
		response.Message = append(response.Message, dbres.Error.Error())
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {

	blogs, db := services.GetAllBlogs(r)
	var status int
	if db.Error == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}

	utils.EncodeJson(w, blogs, status)

}

func GetBlogById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	blogId := vars["id"]
	id, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	blogDetails, dbres := services.GetBlogById(id)
	var status int
	if dbres.Error == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}

	utils.EncodeJson(w, blogDetails, status)
}

func GetBlogsByAuthor(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	username := query.Get("username")
	blogs, db := services.GetBlogsByAuthor(r, username)
	var status int
	if db == nil {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Author doesn't exist")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	} else if db.Error == nil {
		status = http.StatusOK
	}

	utils.EncodeJson(w, blogs, status)

}

func GetSelfBlogs(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	fmt.Println(username)
	blogs, db := services.GetBlogsByAuthor(r, username)
	var status int
	if db == nil {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Author doesn't exist")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	} else if db.Error == nil {
		status = http.StatusOK
	}

	utils.EncodeJson(w, blogs, status)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {

	blog := &models.Blog{}
	utils.ParseBody(r, blog)
	var response dto.Response
	if !services.IsBlogAuthor(r.Header.Get("username"), int(blog.ID)) {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "You are not the author of this blog")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	if blog.ID > 0 {

		blogDetails, dbres := services.GetBlogById(int64(blog.ID))

		if blog.Title != "" {
			blogDetails.Title = blog.Title
		}
		if blog.Body != "" {
			blogDetails.Body = blog.Body
		}

		_, dbres = services.UpdateBlog(blogDetails)
		if dbres.Error == nil {
			response.Status = http.StatusOK
			response.Message = append(response.Message, "Blog successfully updated")
			response.Entity = blogDetails
		} else {
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Failed to update blog")
			response.Message = append(response.Message, dbres.Error.Error())
			response.Entity = nil
		}

		utils.EncodeJson(w, response, response.Status)
	}

}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	blogId := vars["id"]
	id, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	blog, _ := services.GetBlogById(id)
	var response dto.Response
	if !services.IsBlogAuthor(r.Header.Get("username"), int(blog.ID)) {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "You are not the author of this blog")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}
	_, dbres := services.DeleteBlog(id)

	if dbres.Error == nil {
		response.Status = http.StatusOK
		response.Message = append(response.Message, "Blog successfully deleted")
		response.Entity = nil
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to delete blog")
		response.Message = append(response.Message, dbres.Error.Error())
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)
}
