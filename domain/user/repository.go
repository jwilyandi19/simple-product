package user

type UserRepository interface {
	GetAll(req GetUserRequest) ([]User, error)
	Create(req CreateUserRequest) (bool, error)
	GetById(id int) (User, error)
	Update(req UpdateUserRequest) (bool, error)
	Delete(id int) (bool, error)
}
