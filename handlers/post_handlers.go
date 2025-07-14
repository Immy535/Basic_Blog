package handlers

import (
	"blog/models"
	"blog/services"
	"blog/utils"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	Service *services.PostService
}

func (h *PostHandler) ListAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.ListAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) PostByID(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	urlId := v["id"]

	post, err := h.Service.GetPost(urlId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	err = h.Service.CreatePost(&post, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var update models.Post
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	v := mux.Vars(r)
	urlId := v["id"]
	err = h.Service.UpdatePost(&update, urlId, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(update)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	v := mux.Vars(r)
	urlId := v["id"]

	err := h.Service.DeletePost(claims, urlId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}