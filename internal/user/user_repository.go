package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
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
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(lastIdInserted)
	if err != nil {
		return &User{}, err
	}
	user.ID = lastIdInserted
	return user, nil
}
