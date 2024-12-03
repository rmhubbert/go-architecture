package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/mvc/internal/http/controller"
	"github.com/rmhubbert/go-architecture/mvc/internal/repository"
	"github.com/rmhubbert/go-architecture/mvc/internal/service"
)

var dbPath = "../app.db"

func main() {
	roleRepository := repository.NewRoleRepository(dbPath)
	roleService := service.NewRoleService(roleRepository)
	roleController := controller.NewRoleController(roleService)

	userRepository := repository.NewUserRepository(dbPath)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, roleService)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	initRoutes(router, userController, roleController)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
