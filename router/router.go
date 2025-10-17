package router

import (
	"github.com/PatrochR/disko/internal/guild"
	"github.com/PatrochR/disko/internal/user"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine


func InitRouter(userHandler *user.Handler , guildHandler *guild.Handler) {
	r = gin.Default()

	r.POST("/user/register", userHandler.Register)
	r.POST("/user/login", userHandler.Login)

	r.GET("/user/:guildid", userHandler.GetAllUsersByGuildID)

	r.POST("/guild", guildHandler.AddGuild)
	r.GET("/guild/:userid", guildHandler.GetAllGuildsByUserID)

	r.GET("/guild/add", guildHandler.AddUserToGuild)
}

func Start(adder string) error{
	return r.Run(adder)
}
