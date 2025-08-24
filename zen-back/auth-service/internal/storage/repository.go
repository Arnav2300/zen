package storage

import (
	"auth-service/domain"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll(ctx context.Context, page int) ([]*domain.User, int, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	SaveUser(ctx context.Context, user domain.User) error
}

type MagicLinkTokenRepository interface{}
type UserActivationToken interface{}

type Repository struct {
	User UserRepository
}
