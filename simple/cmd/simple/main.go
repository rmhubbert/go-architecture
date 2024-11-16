package main

import (
	"log"
	"net/http"

	"github.com/rmhubbert/go-architecture/simple/internal/app"
)

func main() {
	dbPath := "/Users/hubby/Desktop/test.db"

	userRepository := app.NewUserRepository(dbPath)
	userService := app.NewUserService(userRepository)
	userHandler := app.NewUserHandler(userService)

	roleRepository := app.NewRoleRepository(dbPath)
	roleService := app.NewRoleService(roleRepository)
	roleHandler := app.NewRoleHandler(roleService)

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
