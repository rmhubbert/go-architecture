package dto

import "github.com/rmhubbert/go-architecture/hexagonal/internal/core/domain"

type CreateRoleInput struct {
	Name string `json:"name"`
}

func (cu *CreateRoleInput) Role() *domain.Role {
	return &domain.Role{
		Name: cu.Name,
	}
}

type UpdateRoleInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (cu *UpdateRoleInput) Role() *domain.Role {
	return &domain.Role{
		Id:   cu.Id,
		Name: cu.Name,
	}
}

type RoleOutput struct {
	Id   int    `json:"id,string"`
	Name string `json:"name"`
}

func NewRoleOutput(role *domain.Role) *RoleOutput {
	return &RoleOutput{
		Id:   role.Id,
		Name: role.Name,
	}
}
