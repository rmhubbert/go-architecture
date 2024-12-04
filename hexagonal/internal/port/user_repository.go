package port

import (
	"context"

	"github.com/rmhubbert/go-architecture/hexagonal/internal/core/domain"
)

type UserRepository interface {
	GetOne(ctx context.Context, id int) (*domain.User, error)
	GetMany(ctx context.Context) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id int) error
	UpdatePassword(ctx context.Context, user *domain.User) (*domain.User, error)
}
