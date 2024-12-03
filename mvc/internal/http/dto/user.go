package dto

import "github.com/rmhubbert/go-architecture/mvc/internal/model"

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	RoleId   int    `json:"role_id,string"`
}

func (cu *CreateUserInput) User() *model.User {
	return &model.User{
		Name:     cu.Name,
		Email:    cu.Email,
		Password: cu.Password,
	}
}

type UpdateUserInput struct {
	Id     int    `json:"id,string"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleId int    `json:"role_id,string"`
}

func (cu *UpdateUserInput) User() *model.User {
	return &model.User{
		Id:    cu.Id,
		Name:  cu.Name,
		Email: cu.Email,
	}
}

type UpdateUserPasswordInput struct {
	Id       int    `json:"id,string"`
	Password string `json:"password"`
}

func (pu *UpdateUserPasswordInput) User() *model.User {
	return &model.User{
		Id:       pu.Id,
		Password: pu.Password,
	}
}
