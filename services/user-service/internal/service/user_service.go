package service

import (
	"github.com/prihodkoanton/go-services/services/user-service/internal/domain"
	"github.com/prihodkoanton/go-services/services/user-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *domain.User) error {
	return s.repo.Create(user)
}
