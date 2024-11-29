package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

var dbPath = "../app.db"

func main() {
	roleRepository := app.NewRoleRepository(dbPath)
	roleService := app.NewRoleService(roleRepository)
	roleHandler := app.NewRoleHandler(roleService)

	userRepository := app.NewUserRepository(dbPath)
	userService := app.NewUserService(userRepository)
	userHandler := app.NewUserHandler(userService, roleService)

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
