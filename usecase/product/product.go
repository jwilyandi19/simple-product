package product

import (
	"context"

	domain "github.com/jwilyandi19/simple-product/domain/product"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

type ProductUsecase interface {
	GetProducts(ctx context.Context, req domain.GetProductRequest) ([]domain.Product, error)
	CreateProduct(ctx context.Context, req domain.CreateProductRequest) (bool, error)
}

func NewProductUsecase(p domain.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: p,
	}
}

func (p *productUsecase) GetProducts(ctx context.Context, req domain.GetProductRequest) ([]domain.Product, error) {
	products, err := p.productRepo.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (p *productUsecase) CreateProduct(ctx context.Context, req domain.CreateProductRequest) (bool, error) {
	created, err := p.productRepo.Create(req)
	if err != nil {
		return created, err
	}
	return created, nil
}
