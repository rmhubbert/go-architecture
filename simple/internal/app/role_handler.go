package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	roleService *RoleService
}

func NewRoleHandler(roleService *RoleService) *RoleHandler {
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
	_ = json.NewEncoder(w).Encode(NewRoleOutput(role))
}

func (uh *RoleHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	roles, err := uh.roleService.GetMany(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	outputRoles := []*RoleOutput{}
	for _, role := range roles {
		outputRoles = append(outputRoles, NewRoleOutput(role))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputRoles)
}

func (uh *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	cu := &CreateRoleInput{}
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
	_ = json.NewEncoder(w).Encode(NewRoleOutput(role))
}

func (uh *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cu := &UpdateRoleInput{
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
	_ = json.NewEncoder(w).Encode(NewRoleOutput(role))
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
