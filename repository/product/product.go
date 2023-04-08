package product

import (
	domain "github.com/jwilyandi19/simple-product/domain/product"
	"github.com/jwilyandi19/simple-product/external/db"
)

type productRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}
