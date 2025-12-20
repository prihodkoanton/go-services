package dto

type UserIDUri struct {
	ID int64 `uri:"id" binding:"required,gt=0"`
}
