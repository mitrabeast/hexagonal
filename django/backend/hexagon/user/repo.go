package user

import (
	"context"
)

type UserRepo interface {
	FindByID(ctx context.Context, ID int64) (*User, error)
	Create(ctx context.Context, user *User) error
}
