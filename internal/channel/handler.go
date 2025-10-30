package channel

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddChannel(c *gin.Context) {
	param := c.Param("guild_id")
	var nameReq string
	if err := c.ShouldBindJSON(&nameReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := AddChannelReq{
		Name:    nameReq,
		GuildID: param,
	}

	res, err := h.service.AddGuild(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
