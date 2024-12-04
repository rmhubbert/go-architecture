package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/hexagonal/internal/adapter/http/handler"
	"github.com/rmhubbert/go-architecture/hexagonal/internal/adapter/repository"
	"github.com/rmhubbert/go-architecture/hexagonal/internal/core/service"
)

var dbPath = "../app.db"

func main() {
	roleRepository := repository.NewRoleRepository(dbPath)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	userRepository := repository.NewUserRepository(dbPath)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, roleService)

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	initRoutes(router, userHandler, roleHandler)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
