package user

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	domain "github.com/jwilyandi19/simple-product/domain/user"
	"github.com/jwilyandi19/simple-product/external/cache"
	"github.com/jwilyandi19/simple-product/external/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db    db.SQLDatabase
	cache cache.Cache
}

func NewUserRepository(db db.SQLDatabase, cache cache.Cache) domain.UserRepository {
	return &userRepository{
		db:    db,
		cache: cache,
	}
}

func UserTable() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table("users")
	}
}

func (p *userRepository) GetAll() ([]domain.User, error) {
	var products []domain.User
	db := p.db.Database

	result := db.Scopes(UserTable()).Find(&products)

	if result.Error != nil {
		log.Errorf("[GetAll-User-Repository] %s", result.Error.Error())
		return []domain.User{}, result.Error
	}

	return products, nil
}

func (p *userRepository) Create(req domain.CreateUserRequest) (bool, error) {
	db := p.db.Database
	arg := domain.User{
		FullName: req.FullName,
	}

	err := db.Scopes(UserTable()).Create(&arg).Error
	if err != nil {
		log.Errorf("[Create-User-Repository] %s", err.Error())
		return false, err
	}
	key := fmt.Sprintf("user:%d", arg.ID)
	err = p.setCache(key, arg, p.cache.TTL)
	if err != nil {
		log.Warnf("[Create-User-Repository] redis-set: %s", err.Error())
	}
	return true, nil
}

func (p *userRepository) GetById(id int) (domain.User, error) {
	var user domain.User
	db := p.db.Database

	key := fmt.Sprintf("user:%d", id)
	cachedData, err := p.getCache(key)
	if err != nil {
		log.Warnf("[GetById-User-Repository] redis-get: %s", err.Error())
	} else if cachedData != (domain.User{}) {
		return cachedData, nil
	}

	err = db.Scopes(UserTable()).First(&user, id).Error
	if err != nil {
		log.Errorf("[GetById-User-Repository] %s", err.Error())
		return domain.User{}, err
	}

	err = p.setCache(key, user, p.cache.TTL)
	if err != nil {
		log.Warnf("[GetById-Product-Repository] redis-set: %s", err.Error())
	}
	return user, nil
}

func (p *userRepository) Update(req domain.UpdateUserRequest) (bool, error) {
	db := p.db.Database
	var user domain.User

	err := db.Scopes(UserTable()).First(&user, req.UserID).Error
	if err != nil {
		log.Errorf("[Update-User-Repository] not found %s", err.Error())
		return false, err
	}

	user.ID = req.UserID
	user.FullName = req.FullName

	err = db.Scopes(UserTable()).Save(&user).Error
	if err != nil {
		log.Errorf("[Update-User-Repository] %s", err.Error())
		return false, err
	}

	key := fmt.Sprintf("user:%d", user.ID)
	err = p.setCache(key, user, p.cache.TTL)
	if err != nil {
		log.Warnf("[Update-User-Repository] redis-set: %s", err.Error())
	}

	return true, nil
}

func (p *userRepository) Delete(id int) (bool, error) {
	db := p.db.Database
	var user domain.User

	err := db.Scopes(UserTable()).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		log.Errorf("[Delete-User-Repository] not found %s", err.Error())
		return false, err
	}

	key := fmt.Sprintf("user:%d", id)
	err = p.deleteCache(key)
	if err != nil {
		log.Warnf("[Delete-User-Repository] redis-delete: %s", err.Error())
	}

	return true, nil
}

func (p *userRepository) setCache(key string, data domain.User, ttl int) error {
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

func (p *userRepository) getCache(key string) (domain.User, error) {
	var product domain.User
	conn := p.cache.Redis.Get()
	defer conn.Close()

	cachedData, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if !errors.Is(err, redis.ErrNil) {
			err = errors.New("redis error: " + err.Error())
			return domain.User{}, err
		}
		return domain.User{}, nil
	}
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &product)
		if err != nil {
			return domain.User{}, err
		}
	}
	return product, nil
}

func (p *userRepository) deleteCache(key string) error {
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
