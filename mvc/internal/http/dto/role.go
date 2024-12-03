package dto

import "github.com/rmhubbert/go-architecture/mvc/internal/model"

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
