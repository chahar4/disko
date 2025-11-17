package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetAllUserRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *User) ToGetAllUserRes() GetAllUserRes {
	return GetAllUserRes{
		Username: u.Username,
		Email:    u.Email,
	}
}

type AddUserReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
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
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Repository interface {
	AddUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetAllUsersByGuildID(ctx context.Context, guildID int) (*[]User, error)
}

type Service interface {
	AddUser(ctx context.Context, req *AddUserReq) (*AddUserRes, error)
	Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
	GetAllUsersByGuildID(ctx context.Context, guildID int) (*[]GetAllUserRes, error)
}
