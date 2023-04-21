package user

import (
	"context"
	"encoding/json"

	"github.com/go-playground/validator"
)

type User struct {
	ID        int64  `json:"userId"`
	Username  string `json:"username" validate:"required,gte=5,lte=12"`
	FirstName string `json:"firstName" validate:"required,gte=2,lte=20"`
	LastName  string `json:"lastName" validate:"gte=5,lte=20"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8,lte=50"`
}

func (u *User) PublicInfo() json.RawMessage {
	data, _ := json.Marshal(struct {
		ID        int64  `json:"userId"`
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName,omitempty"`
		Email     string `json:"email"`
	}{
		u.ID, u.Username, u.FirstName, u.LastName, u.Email,
	})
	return data
}

type Users []*User

func (u Users) PublicInfo() json.RawMessage {
	data := make([]json.RawMessage, len(u))
	for i := range u {
		data[i] = u[i].PublicInfo()
	}
	res, _ := json.Marshal(data)
	return res
}

type UserRepo interface {
	FindAll(ctx context.Context) (Users, error)
	FindByID(ctx context.Context, ID int64) (*User, error)
	Create(ctx context.Context, user *User) error
}

type UserService interface {
	List(ctx context.Context) (Users, error)
	Register(ctx context.Context, user *User) error
	Retrieve(ctx context.Context, ID int64) (*User, error)
}

type userService struct {
	repo     UserRepo
	validate *validator.Validate
}

func NewUserService(repo UserRepo) UserService {
	return &userService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *userService) Register(ctx context.Context, user *User) error {
	if err := s.validate.Struct(user); err != nil {
		return err
	}
	return s.repo.Create(ctx, user)
}

func (s *userService) Retrieve(ctx context.Context, ID int64) (*User, error) {
	return s.repo.FindByID(ctx, ID)
}

func (s *userService) List(ctx context.Context) (Users, error) {
	return s.repo.FindAll(ctx)
}
