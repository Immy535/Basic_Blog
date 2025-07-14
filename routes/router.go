package routes

import (
	"blog/handlers"
	"blog/middleware"

	"github.com/gorilla/mux"
)

func SetUpRouter(userHandler *handlers.UserHandler, postHandler *handlers.PostHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register",userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	p := r.PathPrefix("/").Subrouter()
	p.Use(middleware.AuthMiddleware)

	p.HandleFunc("/me", userHandler.LoginInfo).Methods("GET")

	p.HandleFunc("/posts", postHandler.ListAllPosts).Methods("GET")
	p.HandleFunc("/posts/{id}", postHandler.PostByID).Methods("GET")
	p.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	p.HandleFunc("/posts/{id}",postHandler.UpdatePost).Methods("PUT")
	p.HandleFunc("/posts/{id}",postHandler.DeletePost).Methods("DELETE")

	return r
}