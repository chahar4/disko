package guild

import (
	"context"
	"time"
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

func (s *service) AddGuild(ctx context.Context, req *AddGuildReq) (*AddGuildRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	guild := Guild{
		Name:     req.Name,
		OwenerID: req.OwenerID,
	}

	res, err := s.repository.AddGuild(c, &guild)
	if err != nil {
		return nil, err
	}
	return &AddGuildRes{
		ID:       res.ID,
		Name:     res.Name,
		OwenerID: res.OwenerID,
	}, nil
}

func (s *service) GetAllGuildByUserID(ctx context.Context, userID int) (*[]Guild, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res, err := s.repository.GetAllGuildByUserID(c, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) AddUserToGuild(ctx context.Context, req *AddUserToGuildReq) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.repository.AddUserToGuild(c, req.GuildID, req.UserID)
}
