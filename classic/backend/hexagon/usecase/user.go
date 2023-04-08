package usecase

import (
	"context"

	"classic-hexagonal/hexagon/model"
)

type UserUsecase interface {
	Register(ctx context.Context, user *model.User) error
	Retrieve(ctx context.Context, ID int64) (*model.User, error)
	List(ctx context.Context) (model.Users, error)
}
