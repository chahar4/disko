package user

import (
	"net/http"
	"strconv"

	"github.com/PatrochR/disko/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":util.BadRequestErrorMessage.Error()})
		return
	}
	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.AddUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": util.InternalServerErrorMessage.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message:": res.Email + "registred"})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": util.BadRequestErrorMessage.Error()})
		return
	}
	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": util.InternalServerErrorMessage.Error()})
		return
	}

	c.Header("Authorization", res.accessToken)

	r := LoginUserRes{
		ID:       res.ID,
		Username: res.Username,
	}
	c.JSON(http.StatusOK, r)
}

func (h *Handler) GetAllUsersByGuildID(c *gin.Context) {
	guildID_param := c.Param("guild_id")
	guildID, err := strconv.Atoi(guildID_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.GetAllUsersByGuildID(c.Request.Context(), guildID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}
