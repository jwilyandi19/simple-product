package order

import (
	domain "github.com/jwilyandi19/simple-product/domain/order"
	"github.com/jwilyandi19/simple-product/external/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type orderRepository struct {
	db db.SQLDatabase
}

func NewOrderRepository(db db.SQLDatabase) domain.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func OrderTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("order_histories")
	}
}

func (p *orderRepository) GetAll() ([]domain.Order, error) {
	var orders []domain.Order
	db := p.db.Database

	result := db.Scopes(OrderTable()).Find(&orders)

	if result.Error != nil {
		log.Errorf("[GetAll-Order-Repository] %s", result.Error.Error())
		return []domain.Order{}, result.Error
	}

	return orders, nil
}

func (p *orderRepository) Create(req domain.CreateOrderRequest) (bool, error) {
	db := p.db.Database
	arg := domain.Order{
		UserId:       req.UserID,
		OrderItemId:  req.ProductID,
		Descriptions: req.Descriptions,
	}

	err := db.Scopes(OrderTable()).Create(&arg).Error
	if err != nil {
		log.Errorf("[Create-Order-Repository] %s", err.Error())
		return false, err
	}
	return true, nil
}

func (p *orderRepository) GetById(id int) (domain.Order, error) {
	var order domain.Order
	db := p.db.Database

	err := db.Scopes(OrderTable()).First(&order, id).Error
	if err != nil {
		log.Errorf("[GetById-Order-Repository] %s", err.Error())
		return domain.Order{}, err
	}
	return order, nil
}

func (p *orderRepository) Update(req domain.UpdateOrderRequest) (bool, error) {
	db := p.db.Database
	var order domain.Order

	err := db.Scopes(OrderTable()).First(&order, req.OrderID).Error
	if err != nil {
		log.Errorf("[Update-Order-Repository] not found: %s", err.Error())
		return false, err
	}

	order.ID = req.OrderID
	order.UserId = req.UserID
	order.OrderItemId = req.ProductID
	order.Descriptions = req.Descriptions

	err = db.Scopes(OrderTable()).Save(&order).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
