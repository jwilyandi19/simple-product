package product

import (
	"time"

	"gorm.io/gorm"
)

type CreateProductRequest struct {
	ProductName      string
	ProductPrice     int
	ProductExpiredAt time.Time
}

type GetProductRequest struct {
}

type UpdateProductRequest struct {
	ProductID        int
	ProductName      string
	ProductPrice     int
	ProductExpiredAt time.Time
}

type Product struct {
	ID        int            `json:"id" gorm:"primary_key;not null;auto_increment"`
	Name      string         `json:"name" gorm:"type:varchar(1024);"`
	Price     int            `json:"price" gorm:"type:int;"`
	ExpiredAt time.Time      `json:"expired_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
