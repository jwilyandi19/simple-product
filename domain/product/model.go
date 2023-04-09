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
	ExpiredAt time.Time      `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
