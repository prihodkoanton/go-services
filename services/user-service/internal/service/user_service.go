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

func (s *UserService) GetById(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Update(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id int64) error {
	userModel, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	userModel.IsDeleted = true
	return s.repo.Update(userModel)
}
