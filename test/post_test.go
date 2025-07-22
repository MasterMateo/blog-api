package main

import (
	"blog-api/blog"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPostAPI(t *testing.T) {
	store := blog.NewPostStore()
	router := blog.NewRouter(store)

	// Create
	body := `{"title": "Test", "content": "Hello", "author": "Me"}`
	req := httptest.NewRequest("POST", "/posts", bytes.NewReader([]byte(body)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Fatalf("Create failed: %d", resp.Code)
	}

	var created blog.Post
	_ = json.NewDecoder(resp.Body).Decode(&created)

	// Get
	getReq := httptest.NewRequest("GET", "/posts/"+strconv.Itoa(created.ID), nil)
	getResp := httptest.NewRecorder()
	router.ServeHTTP(getResp, getReq)
	if getResp.Code != http.StatusOK {
		t.Fatalf("Get failed: %d", getResp.Code)
	}

	// Update
	update := `{"title": "Updated", "content": "World", "author": "You"}`
	putReq := httptest.NewRequest("PUT", "/posts/"+strconv.Itoa(created.ID), bytes.NewReader([]byte(update)))
	putResp := httptest.NewRecorder()
	router.ServeHTTP(putResp, putReq)
	if putResp.Code != http.StatusOK {
		t.Fatalf("Update failed: %d", putResp.Code)
	}

	// Delete
	delReq := httptest.NewRequest("DELETE", "/posts/"+strconv.Itoa(created.ID), nil)
	delResp := httptest.NewRecorder()
	router.ServeHTTP(delResp, delReq)
	if delResp.Code != http.StatusNoContent {
		t.Fatalf("Delete failed: %d", delResp.Code)
	}

	// Get deleted
	getResp2 := httptest.NewRequest("GET", "/posts/"+strconv.Itoa(created.ID), nil)
	getFail := httptest.NewRecorder()
	router.ServeHTTP(getFail, getResp2)
	if getFail.Code != http.StatusNotFound {
		t.Fatalf("Expected not found after delete, got %d", getFail.Code)
	}
}
