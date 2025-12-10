package domain

import "time"

type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true"`
	FirstName  string    `gorm:"column:first_name"`
	LastName   string    `gorm:"column:last_name"`
	Email      string    `gorm:"column:email;uniqueIndex"`
	Password   string    `gorm:"column:password"`
	CreatedAt  time.Time `gorm:"column:created_at:autoCreateTime"`
	ModifiedAt time.Time `gorm:"column:modified_at:autoUpdateTime"`
	CreatedBy  int64     `gorm:"column:created_by"`
}
