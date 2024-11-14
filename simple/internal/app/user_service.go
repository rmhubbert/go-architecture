package app

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetOne(ctx context.Context, id int) (*User, error) {
	user, err := us.repo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetMany(ctx context.Context) ([]*User, error) {
	users, err := us.repo.GetMany(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) Create(ctx context.Context, user *User) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	user.Password = string(pass)

	user, err = us.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Update(ctx context.Context, user *User) (*User, error) {
	user, err := us.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Delete(ctx context.Context, id int) error {
	err := us.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdatePassword(ctx context.Context, user *User) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	user.Password = string(pass)

	user, err = us.repo.UpdatePassword(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
