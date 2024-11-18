package app

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Roles    []*Role
}

func (user *User) addRole(role *Role) {
	user.Roles = append(user.Roles, role)
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Roles    []int  `json:"roles"`
}

func (cu *CreateUserInput) User() *User {
	return &User{
		Name:     cu.Name,
		Email:    cu.Email,
		Password: cu.Password,
	}
}

type UpdateUserInput struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (cu *UpdateUserInput) User() *User {
	return &User{
		Id:    cu.Id,
		Name:  cu.Name,
		Email: cu.Email,
	}
}

type UserOutput struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserOutput(user *User) *UserOutput {
	return &UserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

type UpdateUserPasswordInput struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

func (pu *UpdateUserPasswordInput) User() *User {
	return &User{
		Id:       pu.Id,
		Password: pu.Password,
	}
}
