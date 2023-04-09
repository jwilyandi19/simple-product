package product

import (
	domain "github.com/jwilyandi19/simple-product/domain/product"
	"github.com/jwilyandi19/simple-product/external/db"
	"gorm.io/gorm"
)

type productRepository struct {
	db db.SQLDatabase
}

func NewProductRepository(db db.SQLDatabase) domain.ProductRepository {
	return &productRepository{
		db: db,
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
		return false, err
	}
	return true, nil
}

func (p *productRepository) GetById(id int) (domain.Product, error) {
	var product domain.Product
	db := p.db.Database

	err := db.Scopes(ProductTable()).First(&product, id).Error
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (p *productRepository) Update(req domain.UpdateProductRequest) (bool, error) {
	db := p.db.Database
	var product domain.Product

	err := db.Scopes(ProductTable()).First(&product, req.ProductID).Error
	if err != nil {
		return false, err
	}

	product.ID = req.ProductID
	product.Name = req.ProductName
	product.Price = req.ProductPrice
	product.ExpiredAt = req.ProductExpiredAt

	err = db.Scopes(ProductTable()).Save(&product).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *productRepository) Delete(id int) (bool, error) {
	db := p.db.Database
	var product domain.Product

	err := db.Scopes(ProductTable()).Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
