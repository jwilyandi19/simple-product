package user

import (
	domain "github.com/jwilyandi19/simple-product/domain/user"
	"github.com/jwilyandi19/simple-product/external/db"
	"gorm.io/gorm"
)

type userRepository struct {
	db db.SQLDatabase
}

func NewUserRepository(db db.SQLDatabase) domain.UserRepository {
	return &userRepository{
		db: db,
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
		return false, err
	}
	return true, nil
}

func (p *userRepository) GetById(id int) (domain.User, error) {
	var user domain.User
	db := p.db.Database

	err := db.Scopes(UserTable()).First(&user, id).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (p *userRepository) Update(req domain.UpdateUserRequest) (bool, error) {
	db := p.db.Database
	var user domain.User

	err := db.Scopes(UserTable()).First(&user, req.UserID).Error
	if err != nil {
		return false, err
	}

	user.ID = req.UserID
	user.FullName = req.FullName

	err = db.Scopes(UserTable()).Save(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *userRepository) Delete(id int) (bool, error) {
	db := p.db.Database
	var user domain.User

	err := db.Scopes(UserTable()).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
