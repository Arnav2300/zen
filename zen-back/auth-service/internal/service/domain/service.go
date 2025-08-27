package domain

import (
	"auth-service/internal/service"
	"auth-service/internal/service/domain/user"
	"auth-service/internal/storage"
)

func NewService(repositories *storage.Repository) *service.Service {
	return &service.Service{
		UserService: user.NewUserService(repositories.User),
	}
}
