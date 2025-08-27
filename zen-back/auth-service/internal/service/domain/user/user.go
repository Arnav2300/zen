package user

import (
	"auth-service/internal/domain"
	"auth-service/internal/storage"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository storage.UserRepository
}

func (s *UserService) GetAllUsers(ctx context.Context, page int) ([]*domain.User, int, error) {
	if page <= 0 {
		return nil, 0, fmt.Errorf("page cannot be <=0")
	}
	users, num, err := s.userRepository.FindAll(ctx, page)
	if err != nil {
		return nil, 0, err
	}
	return users, num, nil

}
func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserService) CreateUser(ctx context.Context, user domain.User) error {
	_, err := s.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("user already exists")
	}

	err = s.userRepository.SaveUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepository storage.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
