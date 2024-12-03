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

type RoleController struct {
	roleService *service.RoleService
}

func NewRoleController(roleService *service.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (uh *RoleController) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		id = 1
	}

	role, err := uh.roleService.GetOne(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(view.NewRoleOutput(role))
}

func (uh *RoleController) GetMany(w http.ResponseWriter, r *http.Request) {
	roles, err := uh.roleService.GetMany(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	outputRoles := []*view.RoleOutput{}
	for _, role := range roles {
		outputRoles = append(outputRoles, view.NewRoleOutput(role))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputRoles)
}

func (uh *RoleController) Create(w http.ResponseWriter, r *http.Request) {
	cu := &dto.CreateRoleInput{}
	err := json.NewDecoder(r.Body).Decode(cu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role, err := uh.roleService.Create(r.Context(), cu.Role())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(view.NewRoleOutput(role))
}

func (uh *RoleController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cu := &dto.UpdateRoleInput{
		Id: id,
	}
	err = json.NewDecoder(r.Body).Decode(cu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role, err := uh.roleService.Update(r.Context(), cu.Role())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(view.NewRoleOutput(role))
}

func (uh *RoleController) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest))
		return
	}

	err = uh.roleService.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
