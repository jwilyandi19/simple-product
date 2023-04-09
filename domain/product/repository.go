package product

type ProductRepository interface {
	GetAll() ([]Product, error)
	Create(req CreateProductRequest) (bool, error)
	GetById(id int) (Product, error)
	Update(req UpdateProductRequest) (bool, error)
	Delete(id int) (bool, error)
}
