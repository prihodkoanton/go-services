package domain

import "time"

type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FirstName  string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName   string    `gorm:"column:last_name;not null" json:"last_name"`
	Email      string    `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Password   string    `gorm:"column:password;not null" json:"password"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	ModifiedAt time.Time `gorm:"column:modified_at;autoUpdateTime" json:"modified_at"`
	CreatedBy  int64     `gorm:"column:created_by" json:"created_by"`
}
