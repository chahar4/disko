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

func (s *service) AddMessage(ctx context.Context, req *AddMessageReq) (*AddMessageRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	message := Message{
		Author_ID:  req.Author_ID,
		Channel_ID: req.Channel_ID,
		Content:    req.Content,
	}

	res, err := s.repository.AddMessage(c, &message)
	if err != nil {
		return nil, err
	}
	return &AddMessageRes{
		ID:         res.ID,
		Author_ID:  res.Author_ID,
		Channel_ID: res.Channel_ID,
		Content:    res.Content,
	}, nil
}
