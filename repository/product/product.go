package product

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	domain "github.com/jwilyandi19/simple-product/domain/product"
	"github.com/jwilyandi19/simple-product/external/cache"
	"github.com/jwilyandi19/simple-product/external/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	db    db.SQLDatabase
	cache cache.Cache
}

func NewProductRepository(db db.SQLDatabase, cache cache.Cache) domain.ProductRepository {
	return &productRepository{
		db:    db,
		cache: cache,
	}
}

func ProductTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("orders_item")
	}
}

func (p *productRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	db := p.db.Database

	result := db.Scopes(ProductTable()).Find(&products)

	if result.Error != nil {
		log.Errorf("[GetAll-Product-Repository] %s", result.Error.Error())
		return []domain.Product{}, result.Error
	}

	return products, nil
}

func (p *productRepository) Create(req domain.CreateProductRequest) (bool, error) {
	db := p.db.Database
	arg := domain.Product{
		Name:      req.ProductName,
		Price:     req.ProductPrice,
		ExpiredAt: req.ProductExpiredAt,
	}

	err := db.Scopes(ProductTable()).Create(&arg).Error
	if err != nil {
		log.Errorf("[Create-Product-Repository] %s", err.Error())
		return false, err
	}
	key := fmt.Sprintf("product:%d", arg.ID)
	err = p.setCache(key, arg, p.cache.TTL)
	if err != nil {
		log.Warnf("[Create-Product-Repository] redis-set: %s", err.Error())
	}
	return true, nil
}

func (p *productRepository) GetById(id int) (domain.Product, error) {
	var product domain.Product
	db := p.db.Database

	key := fmt.Sprintf("product:%d", id)
	cachedData, err := p.getCache(key)
	if err != nil {
		log.Warnf("[GetById-Product-Repository] redis-get: %s", err.Error())
	} else if cachedData != (domain.Product{}) {
		return cachedData, nil
	}

	err = db.Scopes(ProductTable()).First(&product, id).Error
	if err != nil {
		log.Errorf("[GetById-Product-Repository] %s", err.Error())
		return domain.Product{}, err
	}

	err = p.setCache(key, product, p.cache.TTL)
	if err != nil {
		log.Warnf("[GetById-Product-Repository] redis-set: %s", err.Error())
	}
	return product, nil
}

func (p *productRepository) Update(req domain.UpdateProductRequest) (bool, error) {
	db := p.db.Database
	var product domain.Product

	err := db.Scopes(ProductTable()).First(&product, req.ProductID).Error
	if err != nil {
		log.Errorf("[Update-Product-Repository] not found %s", err.Error())
		return false, err
	}

	product.ID = req.ProductID
	product.Name = req.ProductName
	product.Price = req.ProductPrice
	product.ExpiredAt = req.ProductExpiredAt

	err = db.Scopes(ProductTable()).Save(&product).Error
	if err != nil {
		log.Errorf("[Update-Product-Repository] %s", err.Error())
		return false, err
	}

	key := fmt.Sprintf("product:%d", product.ID)
	err = p.setCache(key, product, p.cache.TTL)
	if err != nil {
		log.Warnf("[Update-Product-Repository] redis-set: %s", err.Error())
	}

	return true, nil
}

func (p *productRepository) Delete(id int) (bool, error) {
	db := p.db.Database
	var product domain.Product

	err := db.Scopes(ProductTable()).Where("id = ?", id).Delete(&product).Error
	if err != nil {
		log.Errorf("[Delete-Product-Repository] %s", err.Error())
		return false, err
	}

	key := fmt.Sprintf("product:%d", id)
	err = p.deleteCache(key)
	if err != nil {
		log.Warnf("[Delete-Product-Repository] redis-delete: %s", err.Error())
	}

	return true, nil
}

func (p *productRepository) setCache(key string, data domain.Product, ttl int) error {
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

func (p *productRepository) getCache(key string) (domain.Product, error) {
	var product domain.Product
	conn := p.cache.Redis.Get()
	defer conn.Close()

	cachedData, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if !errors.Is(err, redis.ErrNil) {
			err = errors.New("redis error: " + err.Error())
			return domain.Product{}, err
		}
		return domain.Product{}, nil
	}
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &product)
		if err != nil {
			return domain.Product{}, err
		}
	}
	return product, nil
}

func (p *productRepository) deleteCache(key string) error {
	conn := p.cache.Redis.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("DEL", key))
	if err != nil {
		if !errors.Is(err, redis.ErrNil) {
			err = errors.New("redis error: " + err.Error())
			return err
		}
		return nil
	}
	return nil
}
