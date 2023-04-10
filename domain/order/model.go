package order

import (
	"time"

	"github.com/jwilyandi19/simple-product/domain/product"
	"github.com/jwilyandi19/simple-product/domain/user"
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
	ID           int             `json:"id" gorm:"primary_key;not null;auto_increment"`
	User         user.User       `json:"-" gorm:"foreignkey:user_id"`
	UserId       int             `json:"user_id"`
	OrdersItem   product.Product `json:"-" gorm:"foreignkey:order_item_id"`
	OrderItemId  int             `json:"order_item_id"`
	Descriptions string          `json:"descriptions" gorm:"type:varchar(1024);"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"deleted_at"`
}

type OrderResponse struct {
	ID           int
	UserName     string
	ItemName     string
	Descriptions string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
