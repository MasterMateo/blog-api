package main

// @title Blog API
// @version 1.0
// @description API Documentation
// @host localhost:8080
// @BasePath /

import (
	"log"
	"net/http"

	"blog-api/blog"

	_ "blog-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	store := blog.NewPostStore()
	posts := blog.LoadPostsFromJSON("data/blog_data.json")
	for _, post := range posts {
		store.Create(post)
	}

	r := blog.NewRouter(store)
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler())

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
