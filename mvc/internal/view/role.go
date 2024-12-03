package view

import "github.com/rmhubbert/go-architecture/mvc/internal/model"

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
