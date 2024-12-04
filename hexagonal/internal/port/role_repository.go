package port

import (
	"context"

	"github.com/rmhubbert/go-architecture/hexagonal/internal/core/domain"
)

type RoleRepository interface {
	GetOne(ctx context.Context, id int) (*domain.Role, error)
	GetMany(ctx context.Context) ([]*domain.Role, error)
	GetManyById(ctx context.Context, ids []int) ([]*domain.Role, error)
	Create(ctx context.Context, role *domain.Role) (*domain.Role, error)
	Update(ctx context.Context, role *domain.Role) (*domain.Role, error)
	Delete(ctx context.Context, id int) error
}
