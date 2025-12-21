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

	userModel, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := mapper.ToUserResponse(userModel)
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	var req dto.UserIDUri

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetById(req.ID)
	if err != nil {
		return
	}
	resp := mapper.ToUserResponse(user)
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req dto.UserRequest
	var id dto.UserIDUri

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userModel, err := h.service.Update(&req, &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := mapper.ToUserResponse(userModel)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var req dto.UserEmailUri

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.ToUserResponse(user)
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var req dto.UserIDUri
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.service.Delete(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}
