package main

import (
	"net/http"
	"strconv"
)

type userHandler struct {
	userService *userService
}

func newUserHandler(userService *userService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (uh *userHandler) findOne(w http.ResponseWriter, r *http.Request) {
	user := uh.userService.findOne(r.Context(), 1)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(user.id)))
}
