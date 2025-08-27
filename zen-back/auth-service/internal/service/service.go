package service

import (
	"auth-service/internal/domain"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers(ctx context.Context, page int) ([]*domain.User, int, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}

type Service struct {
	UserService UserService
}
