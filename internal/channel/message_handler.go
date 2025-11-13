package channel

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service     Service
	broadcaster Broadcaster
}

func NewMessageHandler(service Service, broadcaster Broadcaster) *MessageHandler {
	return &MessageHandler{
		service:     service,
		broadcaster: broadcaster,
	}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req AddMessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.broadcaster.BroadcasterMessage([]byte(req.Content), strconv.Itoa(req.Channel_ID))

	//res, err := h.service.AddMessage(c.Request.Context(), &req)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusOK, "d")

}
