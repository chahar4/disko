package user

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

func (h *Handler) Register(c *gin.Context) {
	var req AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.AddUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message:": res.Email + "registred"})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginUserReq
	err := c.ShouldBindJSON(&req)
	if err!= nil{
		c.JSON(http.StatusBadRequest , gin.H{"error": err.Error()})
		return
	}

	res , err := h.service.Login(c.Request.Context(), &req )
	if err!= nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", res.accessToken, 3600, "/", "localhost", false, true)

	r := LoginUserRes{
		ID: res.ID,
		Username: res.Username,
	}
	c.JSON(http.StatusOK, r)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
