package user

import (
	"context"
)

type UserUsecase interface {
	List(ctx context.Context) (Users, error)
	Register(ctx context.Context, user *User) error
	Retrieve(ctx context.Context, ID int64) (*User, error)
}
