package dto

import "github.com/rmhubbert/go-architecture/layered/internal/business/model"

type CreateRoleInput struct {
	Name string `json:"name"`
}

func (cu *CreateRoleInput) Role() *model.Role {
	return &model.Role{
		Name: cu.Name,
	}
}

type UpdateRoleInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (cu *UpdateRoleInput) Role() *model.Role {
	return &model.Role{
		Id:   cu.Id,
		Name: cu.Name,
	}
}

type RoleOutput struct {
	Id   int    `json:"id,string"`
	Name string `json:"name"`
}

func NewRoleOutput(role *model.Role) *RoleOutput {
	return &RoleOutput{
		Id:   role.Id,
		Name: role.Name,
	}
}
