package user

import "github.com/rmhubbert/go-architecture/modular/internal/role"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Role     *role.Role
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	RoleId   int    `json:"role_id,string"`
}

func (cu *CreateUserInput) User() *User {
	return &User{
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

func (cu *UpdateUserInput) User() *User {
	return &User{
		Id:    cu.Id,
		Name:  cu.Name,
		Email: cu.Email,
	}
}

type UserOutput struct {
	Id    int              `json:"id,string"`
	Name  string           `json:"name"`
	Email string           `json:"email"`
	Role  *role.RoleOutput `json:"role"`
}

func NewUserOutput(user *User) *UserOutput {
	return &UserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  role.NewRoleOutput(user.Role),
	}
}

type UpdateUserPasswordInput struct {
	Id       int    `json:"id,string"`
	Password string `json:"password"`
}

func (pu *UpdateUserPasswordInput) User() *User {
	return &User{
		Id:       pu.Id,
		Password: pu.Password,
	}
}
