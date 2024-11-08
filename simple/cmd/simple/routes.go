package main

import (
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

func initRoutes(router *http.ServeMux, userHandler *app.UserHandler) {
	router.HandleFunc("GET /user/{id}", userHandler.Get)
	router.HandleFunc("GET /users", userHandler.GetAll)
	router.HandleFunc("POST /user", userHandler.Create)
	router.HandleFunc("PUT /user", userHandler.Update)
	router.HandleFunc("DELETE /user", userHandler.Delete)
}
