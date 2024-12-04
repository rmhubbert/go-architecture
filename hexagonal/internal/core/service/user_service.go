package service

import (
	"context"

	"github.com/rmhubbert/go-architecture/hexagonal/internal/core/domain"
	"github.com/rmhubbert/go-architecture/hexagonal/internal/port"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetOne(ctx context.Context, id int) (*domain.User, error) {
	user, err := us.repo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetMany(ctx context.Context) ([]*domain.User, error) {
	users, err := us.repo.GetMany(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
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

func (us *UserService) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
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

func (us *UserService) UpdatePassword(
	ctx context.Context,
	user *domain.User,
) (*domain.User, error) {
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
