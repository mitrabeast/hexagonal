package repo

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"django-hexagonal/hexagon/user"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) user.UserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

// Create implements repo.PostgresUserRepo
func (r *PostgresUserRepo) Create(ctx context.Context, user *user.User) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return psql.Insert("users").
		Columns("email", "username", "password", "first_name", "last_name").
		Values(user.Email, user.Username, user.Password, user.FirstName, user.LastName).
		Suffix("RETURNING user_id").
		RunWith(r.db).
		QueryRowContext(ctx).
		Scan(&user.ID)
}

// FindByID implements repo.PostgresUserRepo
func (r *PostgresUserRepo) FindByID(ctx context.Context, ID int64) (*user.User, error) {
	var user user.User
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
