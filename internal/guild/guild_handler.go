package guild

import (
	"net/http"
	"strconv"

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

func (h *Handler) AddGuild(c *gin.Context) {
	var req AddGuildReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.AddGuild(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllGuildsByUserID(c *gin.Context) {
	param := c.Param("userid")

	userID, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guilds, err := h.service.GetAllGuildByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guilds)
}

func (h *Handler) AddUserToGuild(c *gin.Context) {
	userid_param := c.Query("userid")
	guildid_param := c.Query("guildid")
	userID, err := strconv.Atoi(userid_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guildID, err := strconv.Atoi(guildid_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := AddUserToGuildReq{
		UserID:  userID,
		GuildID: guildID,
	}
	if err := h.service.AddUserToGuild(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message: ": "user added to guild"})
}
