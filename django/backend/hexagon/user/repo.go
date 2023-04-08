package user

import (
	"context"
)

type UserRepo interface {
	FindAll(ctx context.Context) (Users, error)
	FindByID(ctx context.Context, ID int64) (*User, error)
	Create(ctx context.Context, user *User) error
}
