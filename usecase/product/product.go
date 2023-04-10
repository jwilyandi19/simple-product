package product

import (
	"context"

	domain "github.com/jwilyandi19/simple-product/domain/product"
	log "github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

type ProductUsecase interface {
	GetProducts(ctx context.Context, req domain.GetProductRequest) ([]domain.Product, error)
	CreateProduct(ctx context.Context, req domain.CreateProductRequest) (bool, error)
	GetDetailProduct(ctx context.Context, id int) (domain.Product, error)
	UpdateProduct(ctx context.Context, req domain.UpdateProductRequest) (bool, error)
	DeleteProduct(ctx context.Context, id int) (bool, error)
}

func NewProductUsecase(p domain.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: p,
	}
}

func (p *productUsecase) GetProducts(ctx context.Context, req domain.GetProductRequest) ([]domain.Product, error) {
	products, err := p.productRepo.GetAll(req)
	if err != nil {
		log.Errorf("[GetProducts-Usecase] %s", err.Error())
		return []domain.Product{}, err
	}
	return products, nil
}

func (p *productUsecase) CreateProduct(ctx context.Context, req domain.CreateProductRequest) (bool, error) {
	created, err := p.productRepo.Create(req)
	if err != nil {
		log.Errorf("[CreateProduct-Usecase] %s", err.Error())
		return created, err
	}
	return created, nil
}

func (p *productUsecase) GetDetailProduct(ctx context.Context, id int) (domain.Product, error) {
	product, err := p.productRepo.GetById(id)
	if err != nil {
		log.Errorf("[GetDetailProduct-Usecase] %s", err.Error())
		return domain.Product{}, err
	}
	return product, nil
}

func (p *productUsecase) UpdateProduct(ctx context.Context, req domain.UpdateProductRequest) (bool, error) {
	updated, err := p.productRepo.Update(req)
	if err != nil {
		log.Errorf("[UpdateProduct-Usecase] %s", err.Error())
		return updated, err
	}
	return updated, nil
}

func (p *productUsecase) DeleteProduct(ctx context.Context, id int) (bool, error) {
	deleted, err := p.productRepo.Delete(id)
	if err != nil {
		log.Errorf("[DeleteProduct-Usecase] %s", err.Error())
		return deleted, err
	}
	return deleted, nil
}
