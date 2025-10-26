package ws

import "github.com/gin-gonic/gin"

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{hub: hub}
}

func (h *Handler) SendMessage(c *gin.Context) {

}
