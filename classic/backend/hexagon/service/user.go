package service

import (
	"context"

	"github.com/go-playground/validator/v10"

	"classic-hexagonal/hexagon/dependency/repo"
	"classic-hexagonal/hexagon/model"
	"classic-hexagonal/hexagon/usecase"
)

type UserService struct {
	repo     repo.UserRepo
	validate *validator.Validate
}

func NewUserService(repo repo.UserRepo) usecase.UserUsecase {
	return &UserService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {
	if err := s.validate.Struct(user); err != nil {
		return err
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Retrieve(ctx context.Context, ID int64) (*model.User, error) {
	return s.repo.FindByID(ctx, ID)
}
