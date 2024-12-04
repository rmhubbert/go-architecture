package dto

import "github.com/rmhubbert/go-architecture/hexagonal/internal/core/domain"

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	RoleId   int    `json:"role_id,string"`
}

func (cu *CreateUserInput) User() *domain.User {
	return &domain.User{
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

func (cu *UpdateUserInput) User() *domain.User {
	return &domain.User{
		Id:    cu.Id,
		Name:  cu.Name,
		Email: cu.Email,
	}
}

type UserOutput struct {
	Id    int         `json:"id,string"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  *RoleOutput `json:"role"`
}

func NewUserOutput(user *domain.User) *UserOutput {
	return &UserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  NewRoleOutput(user.Role),
	}
}

type UpdateUserPasswordInput struct {
	Id       int    `json:"id,string"`
	Password string `json:"password"`
}

func (pu *UpdateUserPasswordInput) User() *domain.User {
	return &domain.User{
		Id:       pu.Id,
		Password: pu.Password,
	}
}
