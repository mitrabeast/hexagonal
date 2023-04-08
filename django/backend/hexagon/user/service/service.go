package service

import (
	"context"

	"github.com/go-playground/validator/v10"

	"django-hexagonal/hexagon/user"
)

type UserService struct {
	repo     user.UserRepo
	validate *validator.Validate
}

func NewUserService(repo user.UserRepo) user.UserUsecase {
	return &UserService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *UserService) Register(ctx context.Context, user *user.User) error {
	if err := s.validate.Struct(user); err != nil {
		return err
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Retrieve(ctx context.Context, ID int64) (*user.User, error) {
	return s.repo.FindByID(ctx, ID)
}

func (s *UserService) List(ctx context.Context) (user.Users, error) {
	return s.repo.FindAll(ctx)
}
