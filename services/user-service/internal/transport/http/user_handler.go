package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prihodkoanton/go-services/services/user-service/internal/dto"
	"github.com/prihodkoanton/go-services/services/user-service/internal/mapper"
	"github.com/prihodkoanton/go-services/services/user-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userModel := mapper.ToUser(&req)
	err := h.service.Create(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := mapper.ToUserResponse(userModel)
	c.JSON(http.StatusOK, resp)
}
