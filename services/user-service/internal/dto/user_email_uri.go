package dto

type UserEmailUri struct {
	Email string `uri:"email" binding:"required,email"`
}
