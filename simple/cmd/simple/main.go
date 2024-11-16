package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

func main() {
	userRepository := app.NewUserRepository("/Users/hubby/Desktop/test.db")
	userService := app.NewUserService(userRepository)
	userHandler := app.NewUserHandler(userService)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	initRoutes(router, userHandler)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
