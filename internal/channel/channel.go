package channel

import "context"

type Channel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GuildID int    `json:"guild_id"`
}

type AddChannelReq struct {
	Name    string `json:"name"`
	GuildID int    `json:"guild_id"`
}

type AddChannelRes struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GuildID int    `json:"guild_id"`
}

type Repository interface {
	AddChannel(ctx context.Context, channel *Channel) (*Channel, error)
}

type Service interface {
	AddGuild(ctx context.Context, req *AddChannelReq) (*AddChannelRes, error)
}
