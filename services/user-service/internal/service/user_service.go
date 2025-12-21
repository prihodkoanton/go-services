package service

import (
	"github.com/prihodkoanton/go-services/services/user-service/internal/domain"
	"github.com/prihodkoanton/go-services/services/user-service/internal/dto"
	"github.com/prihodkoanton/go-services/services/user-service/internal/mapper"
	"github.com/prihodkoanton/go-services/services/user-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(userRequest *dto.UserRequest) (*domain.User, error) {
	user := mapper.ToUser(userRequest)
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetById(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Update(userRequest *dto.UserRequest, id *dto.UserIDUri) (*domain.User, error) {
	userModel, err := s.repo.GetByID(id.ID)
	if err != nil {
		return nil, err
	}
	mapper.ApplyUserUpdate(userModel, userRequest)
	err = s.repo.Update(userModel)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (s *UserService) Delete(id int64) error {
	userModel, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	userModel.IsDeleted = true
	return s.repo.Update(userModel)
}
