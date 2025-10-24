package message

import "context"

type Message struct {
	ID         int    `json:"id"`
	Channel_ID int    `json:"channel_id"`
	Author_ID  int    `json:"author_id"`
	Content    string `json:"content"`
}

type Repository interface {
	AddMessage(ctx context.Context, message *Message) (*Message, error)
}

type Service interface {
	AddMessage(ctx context.Context, req *AddMessageReq) (*AddMessageRes, error)
}

type AddMessageReq struct {
	Channel_ID int    `json:"channel_id"`
	Author_ID  int    `json:"author_id"`
	Content    string `json:"content"`
}

type AddMessageRes struct {
	ID         int    `json:"id"`
	Channel_ID int    `json:"channel_id"`
	Author_ID  int    `json:"author_id"`
	Content    string `json:"content"`
}
