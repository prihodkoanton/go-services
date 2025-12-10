package http

import "github.com/prihodkoanton/go-services/services/user-service/internal/service"

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}

}
