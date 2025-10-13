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

type LoginUserRes struct {
	accessToken string
	ID          string `json:"id"`
	Username    string `json:"username"`
}
type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	AddUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	AddUser(ctx context.Context, req *AddUserReq) (*AddUserRes, error)
	Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
}
