package user

import (
	"time"

	"gorm.io/gorm"
)

type CreateUserRequest struct {
	FullName string
}

type GetUserRequest struct {
}

type UpdateUserRequest struct {
	UserID   int
	FullName string
}

type User struct {
	ID        int            `json:"id" gorm:"primary_key;not null;auto_increment"`
	FullName  string         `json:"full_name" gorm:"type:varchar(1024);"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
