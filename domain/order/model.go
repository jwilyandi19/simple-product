package order

import (
	"time"

	"gorm.io/gorm"
)

type CreateOrderRequest struct {
	UserID       int
	ProductID    int
	Descriptions string
}

type GetOrderRequest struct {
}

type UpdateOrderRequest struct {
	OrderID      int
	UserID       int
	ProductID    int
	Descriptions string
}

type Order struct {
	ID        int            `json:"id" gorm:"primary_key;not null;auto_increment"`
	Name      string         `json:"name" gorm:"type:varchar(1024);"`
	Price     int            `json:"price" gorm:"type:int;"`
	ExpiredAt time.Time      `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
