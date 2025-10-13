package user

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/PatrochR/disko/util"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	repository Repository
	timeout    time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		timeout:    time.Duration(time.Second * 2),
	}
}

func (s *service) AddUser(ctx context.Context, req *AddUserReq) (*AddUserRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashed,
	}

	res, err := s.repository.AddUser(c, &user)
	if err != nil {
		return nil, err
	}
	return &AddUserRes{
		ID:       res.ID,
		Email:    res.Email,
		Password: res.Password,
	}, nil
}



type CustomeClaim struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomeClaim{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	secretKey := os.Getenv("SECRET_KEY")

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{
		accessToken: ss,
		ID:          strconv.Itoa(int(user.ID)),
		Username:    user.Email,
	}, nil
}

