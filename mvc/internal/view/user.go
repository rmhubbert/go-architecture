package view

import "github.com/rmhubbert/go-architecture/mvc/internal/model"

type UserOutput struct {
	Id    int         `json:"id,string"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  *RoleOutput `json:"role"`
}

func NewUserOutput(user *model.User) *UserOutput {
	return &UserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  NewRoleOutput(user.Role),
	}
}
