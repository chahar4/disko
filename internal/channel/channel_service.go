package channel

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

func (s *service) AddGuild(ctx context.Context, req *AddChannelReq) (*AddChannelRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	channel := Channel{
		Name:    req.Name,
		GuildID: req.GuildID,
	}

	res, err := s.repository.AddChannel(c, &channel)
	if err != nil {
		return nil, err
	}
	return &AddChannelRes{
		ID:      res.ID,
		Name:    res.Name,
		GuildID: res.GuildID,
	}, nil
}
