package main

import (
	"blog/database"
	"blog/handlers"
	"blog/repository"
	"blog/routes"
	"blog/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	/* initialise database */
	database.OpenDb()

	/* initialise repositories */
	userRepo := &repository.UserRepo{}
	postRepo := &repository.PostRepo{}

	/* initialise service */
	userService := &services.UserService{Repo: userRepo}
	postService := &services.PostService{Repo: postRepo}

	/* initialise handlers */
	userHandler := &handlers.UserHandler{Service: userService}
	postHandler := &handlers.PostHandler{Service: postService}

	/* initialise routes */
	r := routes.SetUpRouter(userHandler, postHandler)

	/* start server */
	fmt.Println("starting server...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}
