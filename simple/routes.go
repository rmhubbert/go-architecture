package main

import (
	"net/http"
)

func initRoutes(router *http.ServeMux, userHandler *userHandler) {
	router.HandleFunc("GET /user", userHandler.findOne)
}
