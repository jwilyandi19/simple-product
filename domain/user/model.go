package user

import (
	"time"

	"gorm.io/gorm"
)

type CreateUserRequest struct {
	FullName string
}

type GetUserRequest struct {
	Page  int
	Limit int
}

type UpdateUserRequest struct {
	UserID   int
	FullName string
}

type User struct {
	ID        int            `json:"id" gorm:"primary_key;not null;auto_increment"`
	FullName  string         `json:"full_name" gorm:"type:varchar(1024);"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
