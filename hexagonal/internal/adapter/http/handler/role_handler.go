package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rmhubbert/go-architecture/hexagonal/internal/adapter/http/dto"
	"github.com/rmhubbert/go-architecture/hexagonal/internal/port"
)

type RoleHandler struct {
	roleService port.RoleService
}

func NewRoleHandler(roleService port.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (uh *RoleHandler) GetOne(w http.ResponseWriter, r *http.Request) {
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
	_ = json.NewEncoder(w).Encode(dto.NewRoleOutput(role))
}

func (uh *RoleHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	roles, err := uh.roleService.GetMany(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	outputRoles := []*dto.RoleOutput{}
	for _, role := range roles {
		outputRoles = append(outputRoles, dto.NewRoleOutput(role))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputRoles)
}

func (uh *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	_ = json.NewEncoder(w).Encode(dto.NewRoleOutput(role))
}

func (uh *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	_ = json.NewEncoder(w).Encode(dto.NewRoleOutput(role))
}

func (uh *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
