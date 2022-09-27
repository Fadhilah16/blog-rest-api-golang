package controllers

import (
	dto "blog-api-golang/DTO"
	"blog-api-golang/services"
	"blog-api-golang/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	blogId := vars["id"]
	comment := &services.Comment{}
	utils.ParseBody(r, comment)
	author, _ := services.FindUserByUsername(r.Header.Get("username"))

	id, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	blog, dbres := services.GetBlogById(id)
	var response dto.Response
	if dbres.Error != nil {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Cannot comment, blog doesn't exist")
		response.Message = append(response.Message, dbres.Error.Error())
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	_, dbres = services.CreateCommentOnBlog(blog, comment, author)

	if dbres.Error == nil {
		response.Status = http.StatusOK
		response.Message = append(response.Message, "Comment successfully sent")
		response.Entity = comment
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to send comment")
		response.Message = append(response.Message, dbres.Error.Error())
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)

}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment := &services.Comment{}
	utils.ParseBody(r, comment)

	var response dto.Response
	if !services.IsCommentAuthor(r.Header.Get("username"), int(comment.ID)) {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "You are not the author of this comment")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	commentDetails, dbres := services.GetCommentById(int(comment.ID))

	if comment.Body != "" {
		commentDetails.Body = comment.Body
	}

	_, dbres = services.UpdateComment(commentDetails)
	if dbres.Error == nil {
		response.Status = http.StatusOK
		response.Message = append(response.Message, "Comment successfully updated")
		response.Entity = commentDetails
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to update comment")
		response.Message = append(response.Message, dbres.Error.Error())
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["id"]
	id, err := strconv.ParseInt(commentId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	comment, _ := services.GetCommentById(int(id))
	var response dto.Response
	if !services.IsCommentAuthor(r.Header.Get("username"), int(comment.ID)) {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "You are not the author of this comment")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}
	_, dbres := services.DeleteComment(id)

	if dbres.Error == nil {
		response.Status = http.StatusOK
		response.Message = append(response.Message, "Comment successfully deleted")
		response.Entity = nil
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to delete comment")
		response.Message = append(response.Message, dbres.Error.Error())
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)

}
