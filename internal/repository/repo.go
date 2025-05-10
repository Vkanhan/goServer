package repository

import (
	"context"

	"github.com/Vkanhan/newcrud/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.UserCreateDTO) (*domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, id int64, user domain.UserUpdateDTO) (*domain.User, error)
	Delete(ctx context.Context, id int64) error
}