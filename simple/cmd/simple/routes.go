package main

import (
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

func initRoutes(router *http.ServeMux, userHandler *app.UserHandler) {
	router.HandleFunc("GET /user/{id}", userHandler.GetOne)
	router.HandleFunc("GET /users", userHandler.GetMany)
	router.HandleFunc("POST /user", userHandler.Create)
	router.HandleFunc("PATCH /user/{id}", userHandler.Update)
	router.HandleFunc("DELETE /user/{id}", userHandler.Delete)

	router.HandleFunc("PATCH /user/password/{id}", userHandler.UpdatePassword)
}
