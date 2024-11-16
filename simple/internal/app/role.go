package app

type Role struct {
	Id   int
	Name string
}

type CreateRoleInput struct {
	Name string `json:"name"`
}

func (cu *CreateRoleInput) Role() *Role {
	return &Role{
		Name: cu.Name,
	}
}

type UpdateRoleInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (cu *UpdateRoleInput) Role() *Role {
	return &Role{
		Id:   cu.Id,
		Name: cu.Name,
	}
}

type RoleOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewRoleOutput(user *Role) *RoleOutput {
	return &RoleOutput{
		Id:   user.Id,
		Name: user.Name,
	}
}
