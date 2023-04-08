package product

type ProductRepository interface {
	GetAll() ([]Product, error)
	Create(req CreateProductRequest) (bool, error)
}
