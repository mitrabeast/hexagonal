package repo

import (
	"context"

	"classic-hexagonal/hexagon/model"
)

type UserRepo interface {
	FindByID(ctx context.Context, ID int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}
