package blog

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type decodeRequest struct {
	Input string `json:"input"`
}

type decodeResponse struct {
	Result int `json:"result"`
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode JSON response: %v", err)
	}
}

// @Summary Get all posts
// @Tags posts
// @Produce json
// @Success 200 {array} blog.Post
// @Router /posts [get]
func getAllPosts(store *PostStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts := store.GetAll()
		writeJSON(w, http.StatusOK, posts)
	}
}

// @Summary Get post by ID
// @Tags posts
// @Param id path int true "Post ID"
// @Produce json
// @Success 200 {object} blog.Post
// @Failure 404 {string} string "not found"
// @Router /posts/{id} [get]
func getPost(store *PostStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		post, err := store.Get(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, post)
	}
}

// @Summary Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body Post true "New post"
// @Success 201 {object} Post
// @Failure 400 {string} string "bad request"
// @Router /posts [post]
func createPost(store *PostStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if post.Title == "" || post.Content == "" || post.Author == "" {
			http.Error(w, "Missing fields", http.StatusBadRequest)
			return
		}
		created := store.Create(post)
		writeJSON(w, http.StatusCreated, created)
	}
}

// @Summary Update a post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body Post true "Updated post"
// @Success 200 {object} Post
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "not found"
// @Router /posts/{id} [put]
func updatePost(store *PostStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var updated Post
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		post, err := store.Update(id, updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		writeJSON(w, http.StatusOK, post)
	}
}

// @Summary Delete a post
// @Tags posts
// @Param id path int true "Post ID"
// @Success 204 {string} string "deleted"
// @Failure 404 {string} string "not found"
// @Router /posts/{id} [delete]
func deletePost(store *PostStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if err := store.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary Decode numeric message
// @Tags decode
// @Accept json
// @Produce json
// @Param body body decodeRequest true "Input digits"
// @Success 200 {object} decodeResponse
// @Failure 400 {string} string "invalid input"
// @Router /decode [post]
func decodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req decodeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		result := NumDecodings(req.Input)
		writeJSON(w, http.StatusOK, decodeResponse{Result: result})
	}
}
