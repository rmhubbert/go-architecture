package main

import "github.com/rmhubbert/rmhttp"

func initRoutes(rmh *rmhttp.App, userHandler *userHandler) {
	rmh.Get("/user", userHandler.findOne)
}
