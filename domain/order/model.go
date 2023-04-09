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
	User         user.User       `gorm:"foreignkey:user_id"`
	UserId       int             `json:"-"`
	OrdersItem   product.Product `gorm:"foreignkey:order_item_id"`
	OrderItemId  int             `json:"-"`
	Descriptions string          `json:"descriptions" gorm:"type:varchar(1024);"`
	CreatedAt    time.Time       `json:"-"`
	UpdatedAt    time.Time       `json:"-"`
}
