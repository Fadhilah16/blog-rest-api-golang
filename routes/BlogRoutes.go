package routes

import (
	"blog-api-golang/controllers"
	"blog-api-golang/middlewares"

	"github.com/gorilla/mux"
)

var RegisterBlogRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/blogs/", middlewares.Middleware(controllers.CreateBlog)).Methods("POST")
	router.HandleFunc("/api/blogs/", controllers.GetBlogs).Methods("GET")
	router.HandleFunc("/api/blogs/{id}", controllers.GetBlogById).Methods("GET")
	router.HandleFunc("/api/blogs/u/me", middlewares.Middleware(controllers.GetSelfBlogs)).Methods("GET")
	router.HandleFunc("/api/blogs/u/", controllers.GetBlogsByAuthor).Methods("GET")
	router.HandleFunc("/api/blogs/", middlewares.Middleware(controllers.UpdateBlog)).Methods("PUT")
	router.HandleFunc("/api/blogs/{id}", middlewares.Middleware(controllers.DeleteBlog)).Methods("DELETE")

}
