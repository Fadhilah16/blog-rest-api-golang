package routes

import (
	"blog-api-golang/controllers"
	"blog-api-golang/middlewares"

	"github.com/gorilla/mux"
)

var RegisterCommentRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/blogs/{id}/comment", middlewares.Middleware(controllers.CreateComment)).Methods("POST")
	router.HandleFunc("/api/blogs/comment/update", middlewares.Middleware(controllers.UpdateComment)).Methods("PUT")
	router.HandleFunc("/api/blogs/comment/delete/{id}", middlewares.Middleware(controllers.DeleteComment)).Methods("DELETE")
}
