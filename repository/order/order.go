package order

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	domain "github.com/jwilyandi19/simple-product/domain/order"
	"github.com/jwilyandi19/simple-product/external/cache"
	"github.com/jwilyandi19/simple-product/external/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type orderRepository struct {
	db    db.SQLDatabase
	cache cache.Cache
}

func NewOrderRepository(db db.SQLDatabase, cache cache.Cache) domain.OrderRepository {
	return &orderRepository{
		db:    db,
		cache: cache,
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
	key := fmt.Sprintf("order:%d", arg.ID)
	err = p.setCache(key, arg, p.cache.TTL)
	if err != nil {
		log.Warnf("[Create-Order-Repository] redis-set: %s", err.Error())
	}
	return true, nil
}

func (p *orderRepository) GetById(id int) (domain.Order, error) {
	var order domain.Order
	db := p.db.Database

	key := fmt.Sprintf("order:%d", id)
	cachedData, err := p.getCache(key)
	if err != nil {
		log.Warnf("[GetById-Order-Repository] redis-get: %s", err.Error())
	} else if cachedData != (domain.Order{}) {
		return cachedData, nil
	}

	err = db.Scopes(OrderTable()).First(&order, id).Error
	if err != nil {
		log.Errorf("[GetById-Order-Repository] %s", err.Error())
		return domain.Order{}, err
	}

	err = p.setCache(key, order, p.cache.TTL)
	if err != nil {
		log.Warnf("[GetById-Order-Repository] redis-set: %s", err.Error())
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

	key := fmt.Sprintf("order:%d", order.ID)
	err = p.setCache(key, order, p.cache.TTL)
	if err != nil {
		log.Warnf("[Update-Order-Repository] redis-set: %s", err.Error())
	}

	return true, nil
}

func (p *orderRepository) setCache(key string, data domain.Order, ttl int) error {
	conn := p.cache.Redis.Get()
	defer conn.Close()

	marshalledData, err := json.Marshal(&data)
	if err != nil {
		return errors.New("marshal error: " + err.Error())
	}
	_, err = conn.Do("SETEX", key, ttl, string(marshalledData))
	if err != nil {
		return errors.New("redis error: " + err.Error())
	}

	return nil
}

func (p *orderRepository) getCache(key string) (domain.Order, error) {
	var product domain.Order
	conn := p.cache.Redis.Get()
	defer conn.Close()

	cachedData, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if !errors.Is(err, redis.ErrNil) {
			err = errors.New("redis error: " + err.Error())
			return domain.Order{}, err
		}
		return domain.Order{}, nil
	}
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &product)
		if err != nil {
			return domain.Order{}, err
		}
	}
	return product, nil
}
