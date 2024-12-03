package main

import (
	"net/http"

	"github.com/rmhubbert/go-architecture/mvc/internal/http/controller"
)

func initRoutes(
	router *http.ServeMux,
	userController *controller.UserController,
	roleController *controller.RoleController,
) {
	router.HandleFunc("GET /user/{id}", userController.GetOne)
	router.HandleFunc("GET /users", userController.GetMany)
	router.HandleFunc("POST /user", userController.Create)
	router.HandleFunc("PATCH /user/{id}", userController.Update)
	router.HandleFunc("DELETE /user/{id}", userController.Delete)
	router.HandleFunc("PATCH /user/{id}/password", userController.UpdatePassword)

	router.HandleFunc("GET /role/{id}", roleController.GetOne)
	router.HandleFunc("GET /roles", roleController.GetMany)
	router.HandleFunc("POST /role", roleController.Create)
	router.HandleFunc("PATCH /role/{id}", roleController.Update)
	router.HandleFunc("DELETE /role/{id}", roleController.Delete)
}
