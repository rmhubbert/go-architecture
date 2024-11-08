package app

import "context"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Get(ctx context.Context, id int) (*User, error) {
	return &User{
		Id:    id,
		Name:  "Hubby",
		Email: "hubby@rmhubbert.com",
	}, nil
}

func (us *UserService) GetAll(ctx context.Context) ([]*User, error) {
	user1 := &User{
		Id:    1,
		Name:  "Hubby",
		Email: "hubby@rmhubbert.com",
	}
	user2 := &User{
		Id:    2,
		Name:  "Hubby2",
		Email: "hubby2@rmhubbert.com",
	}

	return []*User{user1, user2}, nil
}

func (us *UserService) Create(ctx context.Context, user *User) (*User, error) {
	// TODO: insert user
	return user, nil
}

func (us *UserService) Update(ctx context.Context, user *User) (*User, error) {
	// TODO: update user
	return user, nil
}

func (us *UserService) Delete(ctx context.Context, user *User) (*User, error) {
	// TODO: delete user
	return user, nil
}
