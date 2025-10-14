package router

import (
	"github.com/PatrochR/disko/internal/user"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine


func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/user/register", userHandler.Register)
	r.POST("/user/login", userHandler.Login)
}

func Start(adder string) error{
	return r.Run(adder)
}
