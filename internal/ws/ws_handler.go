package ws

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PatrochR/disko/internal/channel"
	"github.com/PatrochR/disko/internal/guild"
	"github.com/PatrochR/disko/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

type Handler struct {
	guildService   guild.Service
	channelService channel.Service
	hub            *Hub
}

func NewHandler(hub *Hub, guildService guild.Service, channelService channel.Service) *Handler {
	return &Handler{
		hub:            hub,
		guildService:   guildService,
		channelService: channelService,
	}
}

var upgrade = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) OpenWS(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return
	}

	query := c.Query("token")

	trimHeader := strings.TrimPrefix(query, "Bearer ")

	secretKey := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(trimHeader, &user.CustomeClaim{}, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error ": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*user.CustomeClaim)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error ": "Unauthorized"})
		c.Abort()
		return
	}

	userIDInt, err := strconv.Atoi(claims.ID)
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return
	}

	guildIDs, err := h.guildService.GetAllGuildByUserID(c.Request.Context(), userIDInt)
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return
	}

	os.Stdout.Sync()
	channelIDs := make(map[string]struct{})

	for _, guildID := range *guildIDs {
		channels, err := h.channelService.GetChannelsByGuildID(c.Request.Context(), guildID.ID)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			return
		}
		for _, channel := range *channels {
			channelIDs[strconv.Itoa(channel.ID)] = struct{}{}
		}

	}

	client := Client{
		Send:  make(chan []byte, 10),
		Conn:  conn,
		Rooms: channelIDs,
		hub:   h.hub,
	}

	h.hub.Register <- &client
	client.WriteMessage()
}
