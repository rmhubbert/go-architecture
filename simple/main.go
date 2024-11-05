package main

import (
	"log"

	"github.com/rmhubbert/rmhttp"
)

func main() {
	userService := newUserService()
	userHandler := newUserHandler(userService)

	rmh := rmhttp.New()
	initRoutes(rmh, userHandler)

	log.Fatal(rmh.Start())
}
