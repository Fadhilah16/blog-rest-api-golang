package main

import (
	"fmt"
	"log"
	"net/http"

	"blog-api-golang/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBlogRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterCommentRoutes(r)
	http.Handle("/", r)

	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
