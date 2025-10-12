package user

import (
	"context"
	"time"

	"github.com/PatrochR/disko/util"
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
