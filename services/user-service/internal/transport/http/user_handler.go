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

// @Summary Create new user
// @Description Create new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserRequest true "User data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userModel, err := h.service.Create(&req)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponse(userModel))
}

// @Summary Get user by id
// @Description Get user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	var req dto.UserIDUri
	if err := c.ShouldBindUri(&req); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.GetById(req.ID)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponse(user))
}

// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Param user body dto.UserRequest true "User data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req dto.UserRequest
	var id dto.UserIDUri

	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindUri(&id); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userModel, err := h.service.Update(&req, &id)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponse(userModel))
}

// @Summary Get user by email
// @Description Get user by email
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var req dto.UserEmailUri
	if err := c.ShouldBindUri(&req); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.GetByEmail(req.Email)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, mapper.ToUserResponse(user))
}

// @Summary Delete user
// @Description Delete user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 202
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	var req dto.UserIDUri
	if err := c.ShouldBindUri(&req); err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Delete(req.ID); err != nil {
		respondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusAccepted)
}

func respondError(c *gin.Context, code int, msg string) {
	c.JSON(code, dto.ErrorResponse{
		Code:    code,
		Message: msg,
	})
}
