package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserRes struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	AddUser(ctx context.Context, user *User) (*User, error)
}

type Service interface {
	AddUser(ctx context.Context, req *AddUserReq) (*AddUserRes, error)
}
