package handler

import "github.com/jwilyandi19/simple-product/usecase/product"

type productHandler struct {
	productUsecase product.ProductUsecase
}

type ProductHandler interface {
}

func NewProductHandler(product product.ProductUsecase) ProductHandler {
	return &productHandler{
		productUsecase: product,
	}
}
