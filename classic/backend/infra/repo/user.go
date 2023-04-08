package repo

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"classic-hexagonal/hexagon/dependency/repo"
	"classic-hexagonal/hexagon/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Create implements repo.UserRepo
func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return psql.Insert("users").
		Columns("email", "username", "password", "first_name", "last_name").
		Values(user.Email, user.Username, user.Password, user.FirstName, user.LastName).
		Suffix("RETURNING user_id").
		RunWith(r.db).
		QueryRowContext(ctx).
		Scan(&user.ID)
}

// FindByID implements repo.UserRepo
func (r *UserRepo) FindByID(ctx context.Context, ID int64) (*model.User, error) {
	var user model.User
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	if err := psql.Select("user_id", "email", "username", "password", "first_name", "last_name").
		From("users").
		Where(sq.Eq{"user_id": ID}).
		RunWith(r.db).
		QueryRowContext(ctx).
		Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}
	return &user, nil
}
