package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

func main() {
	userService := app.NewUserService()
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
