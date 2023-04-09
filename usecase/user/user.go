package user

import (
	"context"

	domain "github.com/jwilyandi19/simple-product/domain/user"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

type UserUsecase interface {
	GetUsers(ctx context.Context, req domain.GetUserRequest) ([]domain.User, error)
	CreateUser(ctx context.Context, req domain.CreateUserRequest) (bool, error)
	GetDetailUser(ctx context.Context, id int) (domain.User, error)
	UpdateUser(ctx context.Context, req domain.UpdateUserRequest) (bool, error)
	DeleteUser(ctx context.Context, id int) (bool, error)
}

func NewUserUsecase(p domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: p,
	}
}

func (p *userUsecase) GetUsers(ctx context.Context, req domain.GetUserRequest) ([]domain.User, error) {
	users, err := p.userRepo.GetAll()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (p *userUsecase) CreateUser(ctx context.Context, req domain.CreateUserRequest) (bool, error) {
	created, err := p.userRepo.Create(req)
	if err != nil {
		return created, err
	}
	return created, nil
}

func (p *userUsecase) GetDetailUser(ctx context.Context, id int) (domain.User, error) {
	user, err := p.userRepo.GetById(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (p *userUsecase) UpdateUser(ctx context.Context, req domain.UpdateUserRequest) (bool, error) {
	updated, err := p.userRepo.Update(req)
	if err != nil {
		return updated, err
	}
	return updated, nil
}

func (p *userUsecase) DeleteUser(ctx context.Context, id int) (bool, error) {
	deleted, err := p.userRepo.Delete(id)
	if err != nil {
		return deleted, err
	}
	return deleted, nil
}
