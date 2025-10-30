package router

import (
	"github.com/PatrochR/disko/internal/channel"
	"github.com/PatrochR/disko/internal/guild"
	"github.com/PatrochR/disko/internal/user"
	"github.com/PatrochR/disko/middleware"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(
	userHandler *user.Handler,
	guildHandler *guild.Handler,
	channelHandler *channel.Handler) {
	r = gin.Default()

	api :=r.Group("/api/v1")

	api.POST("/user/register", userHandler.Register)
	api.POST("/user/login", userHandler.Login)

	authorized := api.Group("/")
	authorized.Use(middleware.JwtAuth())


	authorized.GET("/user/:guildid", userHandler.GetAllUsersByGuildID)
	authorized.POST("/guild", guildHandler.AddGuild)
	authorized.GET("/guild/:userid", guildHandler.GetAllGuildsByUserID)
	authorized.GET("/guild/add", guildHandler.AddUserToGuild)
	authorized.POST("/channel", channelHandler.AddChannel)
}

func Start(adder string) error {
	return r.Run(adder)
}
