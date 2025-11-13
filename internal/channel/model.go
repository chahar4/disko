package channel

import "context"

// channel
type Channel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GuildID int    `json:"guild_id"`
}

type AddChannelReq struct {
	Name    string `json:"name"`
	GuildID string `json:"guild_id"`
}

type AddChannelRes struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GuildID int    `json:"guild_id"`
}

// message
type AddMessageReq struct {
	Channel_ID int    `json:"channel_id"`
	Author_ID  int    `json:"author_id"`
	Content    string `json:"content"`
}

type Message struct {
	ID         int    `json:"id"`
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

// interfaces
type Repository interface {
	GetChannelsByGuildID(ctx context.Context, guildID int) (*[]Channel, error)
	AddChannel(ctx context.Context, channel *Channel) (*Channel, error)
	AddMessage(ctx context.Context, message *Message) (*Message, error)
}

type Service interface {
	GetChannelsByGuildID(ctx context.Context, guildID int) (*[]Channel, error)
	AddChannel(ctx context.Context, req *AddChannelReq) (*AddChannelRes, error)
	AddMessage(ctx context.Context, req *AddMessageReq) (*AddMessageRes, error)
}

type Broadcaster interface {
	BroadcasterMessage(payload []byte, roomID string)
}
