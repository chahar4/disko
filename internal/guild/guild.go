package guild

import "context"

type Guild struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	OwenerID int    `json:"owner_id"`
}

type Repository interface {
	AddGuild(ctx context.Context, guild *Guild) (*Guild, error)
	GetAllGuildByUserID(ctx context.Context, userID int) (*[]Guild, error)
	AddUserToGuild(ctx context.Context, guildID, userID int) error
}

type Service interface {
	AddGuild(ctx context.Context, req *AddGuildReq) (*AddGuildRes, error)
	GetAllGuildByUserID(ctx context.Context, userID int) (*[]Guild, error)
	AddUserToGuild(ctx context.Context, req *AddUserToGuildReq) error
}

type AddGuildRes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	OwenerID int    `json:"owner_id"`
}

type AddGuildReq struct {
	Name     string `json:"name"`
	OwenerID int    `json:"owner_id"`
}

type AddUserToGuildReq struct {
	UserID  int `json:"user_id" uri:"userid" binding:"required"`
	GuildID int `json:"guild_id" uri:"guildid" binding:"required"`
}
