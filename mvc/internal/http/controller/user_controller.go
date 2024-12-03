package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rmhubbert/go-architecture/mvc/internal/http/dto"
	"github.com/rmhubbert/go-architecture/mvc/internal/service"
	"github.com/rmhubbert/go-architecture/mvc/internal/view"
)

type UserController struct {
	userService *service.UserService
	roleService *service.RoleService
}

func NewUserController(
	userService *service.UserService,
	roleService *service.RoleService,
) *UserController {
	return &UserController{
		userService: userService,
		roleService: roleService,
	}
}

func (uh *UserController) GetOne(w http.ResponseWriter, r *http.Request) {
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
	_ = json.NewEncoder(w).Encode(view.NewUserOutput(user))
}

func (uh *UserController) GetMany(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.GetMany(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	outputUsers := []*view.UserOutput{}
	for _, user := range users {
		outputUsers = append(outputUsers, view.NewUserOutput(user))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputUsers)
}

func (uh *UserController) Create(w http.ResponseWriter, r *http.Request) {
	cu := &dto.CreateUserInput{}
	err := json.NewDecoder(r.Body).Decode(cu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role, err := uh.roleService.GetOne(r.Context(), cu.RoleId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	user := cu.User()
	user.Role = role

	user, err = uh.userService.Create(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(view.NewUserOutput(user))
}

func (uh *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cu := &dto.UpdateUserInput{
		Id: id,
	}
	err = json.NewDecoder(r.Body).Decode(cu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role, err := uh.roleService.GetOne(r.Context(), cu.RoleId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	user := cu.User()
	user.Role = role

	user, err = uh.userService.Update(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(view.NewUserOutput(user))
}

func (uh *UserController) Delete(w http.ResponseWriter, r *http.Request) {
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

func (uh *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pu := &dto.UpdateUserPasswordInput{
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
	_ = json.NewEncoder(w).Encode(view.NewUserOutput(user))
}
