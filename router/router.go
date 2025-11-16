package router

import (
	"github.com/PatrochR/disko/internal/channel"
	"github.com/PatrochR/disko/internal/guild"
	"github.com/PatrochR/disko/internal/user"
	"github.com/PatrochR/disko/internal/ws"
	"github.com/PatrochR/disko/middleware"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(
	userHandler *user.Handler,
	guildHandler *guild.Handler,
	channelHandler *channel.Handler,
	wsHandler *ws.Handler) {
	r = gin.Default()

	api := r.Group("/api/v1")

	api.POST("/auth/register", userHandler.Register)
	api.POST("/auth/login", userHandler.Login)

	api.GET("/ws", wsHandler.OpenWS)

	authorized := api.Group("/")
	authorized.Use(middleware.JwtAuth())

	authorized.GET("/guilds/:guild_id/users", userHandler.GetAllUsersByGuildID)
	authorized.POST("/guilds", guildHandler.AddGuild)
	authorized.GET("/users/:user_id/guilds", guildHandler.GetAllGuildsByUserID)
	authorized.GET("/guilds/:guild_id/members", guildHandler.AddUserToGuild)
	authorized.POST("/guilds/:guild_id/channels", channelHandler.AddChannel)

	api.POST("/message", channelHandler.SendMessage)

}

func Start(adder string) error {
	return r.Run(adder)
}
