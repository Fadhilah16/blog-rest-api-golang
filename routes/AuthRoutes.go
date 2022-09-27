package routes

import (
	"blog-api-golang/controllers"
	"blog-api-golang/middlewares"

	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/auth/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/auth/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/api/auth/account/update", middlewares.Middleware(controllers.UpdateProfile)).Methods("POST")
}
