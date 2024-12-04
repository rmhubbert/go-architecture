package service

import (
	"context"

	"github.com/rmhubbert/go-architecture/hexagonal/internal/core/domain"
	"github.com/rmhubbert/go-architecture/hexagonal/internal/port"
)

type RoleService struct {
	repo port.RoleRepository
}

func NewRoleService(repo port.RoleRepository) *RoleService {
	return &RoleService{
		repo: repo,
	}
}

func (rs *RoleService) GetOne(ctx context.Context, id int) (*domain.Role, error) {
	role, err := rs.repo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (rs *RoleService) GetMany(ctx context.Context) ([]*domain.Role, error) {
	roles, err := rs.repo.GetMany(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (rs *RoleService) GetManyById(ctx context.Context, ids []int) ([]*domain.Role, error) {
	roles, err := rs.repo.GetManyById(ctx, ids)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (rs *RoleService) Create(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	role, err := rs.repo.Create(ctx, role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (rs *RoleService) Update(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	role, err := rs.repo.Update(ctx, role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (rs *RoleService) Delete(ctx context.Context, id int) error {
	err := rs.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
