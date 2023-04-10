package order

type OrderRepository interface {
	GetAll(req GetOrderRequest) ([]Order, error)
	Create(req CreateOrderRequest) (bool, error)
	GetById(id int) (Order, error)
	Update(req UpdateOrderRequest) (bool, error)
}
