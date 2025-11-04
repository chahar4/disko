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
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ownerID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ownerIDInt, err := strconv.Atoi(ownerID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	addGuildReq := AddGuildReq{
		Name:     req.Name,
		OwenerID: ownerIDInt,
	}

	res, err := h.service.AddGuild(c.Request.Context(), &addGuildReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllGuildsByUserID(c *gin.Context) {
	param := c.Param("user_id")

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
	guild_param := c.Param("guild_id")
	userid_param := c.Query("user_id")
	userID, err := strconv.Atoi(userid_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guildID, err := strconv.Atoi(guild_param)
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
