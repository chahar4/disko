package router

import (
	"github.com/PatrochR/disko/internal/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Adder       string
	Engine      *gin.Engine
	userHandler *user.Handler
}

func NewRouter(Adder string, Engine *gin.Engine, userHandler *user.Handler) *Router {
	return &Router{
		Adder:       Adder,
		Engine:      Engine,
		userHandler: userHandler,
	}
}

func (r *Router) Start() error {
	r.Engine.POST("/user/register", r.userHandler.Register)
	r.Engine.POST("/user/login", r.userHandler.Login)

	return r.Engine.Run(r.Adder)
}
