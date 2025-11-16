package channel

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
	broadcaster Broadcaster
}

func NewHandler(service Service , broadcaster Broadcaster) *Handler {
	return &Handler{
		service: service,
		broadcaster: broadcaster,
	}
}

func (h *Handler) AddChannel(c *gin.Context) {
	param := c.Param("guild_id")
	var reqJSON struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&reqJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := AddChannelReq{
		Name:    reqJSON.Name,
		GuildID: param,
	}

	res, err := h.service.AddChannel(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) SendMessage(c *gin.Context) {
	var req AddMessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.broadcaster.BroadcasterMessage([]byte(req.Content), strconv.Itoa(req.Channel_ID))

	_, err := h.service.AddMessage(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "send")

}
