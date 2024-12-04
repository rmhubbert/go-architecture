package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/layered/internal/business/service"
	"github.com/rmhubbert/go-architecture/layered/internal/presentation/http/handler"
	"github.com/rmhubbert/go-architecture/layered/internal/store"
)

var dbPath = "../app.db"

func main() {
	roleRepository := store.NewRoleRepository(dbPath)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	userRepository := store.NewUserRepository(dbPath)
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
