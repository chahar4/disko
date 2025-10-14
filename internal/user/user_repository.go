package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users(username , email , password) VALUES ($1 , $2 , $3) returning id"

	var lastIdInserted int64
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastIdInserted)
	if err != nil {
		return &User{}, err
	}
	user.ID = lastIdInserted
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {

	query := "SELECT * FROM users WHERE email = $1"
	var user User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
