package main

import (
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

func initRoutes(router *http.ServeMux, userHandler *app.UserHandler, roleHandler *app.RoleHandler) {
	router.HandleFunc("GET /user/{id}", userHandler.GetOne)
	router.HandleFunc("GET /users", userHandler.GetMany)
	router.HandleFunc("POST /user", userHandler.Create)
	router.HandleFunc("PATCH /user/{id}", userHandler.Update)
	router.HandleFunc("DELETE /user/{id}", userHandler.Delete)
	router.HandleFunc("PATCH /user/{id}/password", userHandler.UpdatePassword)

	router.HandleFunc("GET /role/{id}", roleHandler.GetOne)
	router.HandleFunc("GET /roles", roleHandler.GetMany)
	router.HandleFunc("POST /role", roleHandler.Create)
	router.HandleFunc("PATCH /role/{id}", roleHandler.Update)
	router.HandleFunc("DELETE /role/{id}", roleHandler.Delete)
}
