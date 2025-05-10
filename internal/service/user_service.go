package service

import (
	"context"
	"errors"
	"time"

	"github.com/Vkanhan/newcrud/internal/domain"
	"github.com/Vkanhan/newcrud/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, dto domain.UserCreateDTO) (*domain.User, error) {
	if dto.Email == "" {
		return nil, errors.New("email is required")
	}
	if dto.FirstName == "" {
		return nil, errors.New("first name is required")
	}

	// Business logic can go here
	// checking for duplicate emails, formatting data, etc.

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.Create(ctx, dto)
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.GetAll(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, dto domain.UserUpdateDTO) (*domain.User, error) {
	if dto.Email == "" {
		return nil, errors.New("email is required")
	}
	if dto.FirstName == "" {
		return nil, errors.New("first name is required")
	}

	_, err := s.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.Update(ctx, id, dto)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.GetUserByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.Delete(ctx, id)
}
