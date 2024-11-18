package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		id = 1
	}

	user, err := uh.userService.GetOne(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(NewUserOutput(user))
}

func (uh *UserHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.GetMany(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	outputUsers := []*UserOutput{}
	for _, user := range users {
		outputUsers = append(outputUsers, NewUserOutput(user))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputUsers)
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	cu := &CreateUserInput{}
	err := json.NewDecoder(r.Body).Decode(cu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uh.userService.Create(r.Context(), cu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(NewUserOutput(user))
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cu := &UpdateUserInput{
		Id: id,
	}
	err = json.NewDecoder(r.Body).Decode(cu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uh.userService.Update(r.Context(), cu.User())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(NewUserOutput(user))
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}

	err = uh.userService.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pu := &UpdateUserPasswordInput{
		Id: id,
	}
	err = json.NewDecoder(r.Body).Decode(pu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uh.userService.UpdatePassword(r.Context(), pu.User())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(NewUserOutput(user))
}
