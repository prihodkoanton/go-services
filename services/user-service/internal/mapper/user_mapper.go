package mapper

import (
	"github.com/prihodkoanton/go-services/services/user-service/internal/domain"
	"github.com/prihodkoanton/go-services/services/user-service/internal/dto"
)

func ToUser(r *dto.UserRequest) *domain.User {
	return &domain.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
		CreatedBy: -1,
	}
}

func ToUserResponse(u *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
