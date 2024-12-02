package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/modular/internal/role"
	"github.com/rmhubbert/go-architecture/modular/internal/user"
)

var dbPath = "../app.db"

func main() {
	roleRepository := role.NewRoleRepository(dbPath)
	roleService := role.NewRoleService(roleRepository)
	roleHandler := role.NewRoleHandler(roleService)

	userRepository := user.NewUserRepository(dbPath)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService, roleService)

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
