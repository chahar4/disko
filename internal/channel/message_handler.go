package channel

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service Service
}

func NewmessageHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SendMessage(c *gin.Context) {
	var req AddMessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.AddMessage(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
