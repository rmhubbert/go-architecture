package main

import (
	"log"
	"net/http"
)

func main() {
	userService := newUserService()
	userHandler := newUserHandler(userService)

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
