package blog

import (
	"github.com/gorilla/mux"
)

func NewRouter(store *PostStore) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/posts", getAllPosts(store)).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", getPost(store)).Methods("GET")
	r.HandleFunc("/posts", createPost(store)).Methods("POST")
	r.HandleFunc("/posts/{id:[0-9]+}", updatePost(store)).Methods("PUT")
	r.HandleFunc("/posts/{id:[0-9]+}", deletePost(store)).Methods("DELETE")
	r.HandleFunc("/decode", decodeHandler()).Methods("POST")
	return r
}
