package product

import (
	domain "github.com/jwilyandi19/simple-product/domain/product"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

type ProductUsecase interface {
}

func NewProductUsecase(p domain.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: p,
	}
}
