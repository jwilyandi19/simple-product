package order

import (
	"context"
	"sync"

	domain "github.com/jwilyandi19/simple-product/domain/order"
	"github.com/jwilyandi19/simple-product/domain/product"
	"github.com/jwilyandi19/simple-product/domain/user"
	log "github.com/sirupsen/logrus"
)

type orderUsecase struct {
	orderRepo   domain.OrderRepository
	productRepo product.ProductRepository
	userRepo    user.UserRepository
}

type OrderUsecase interface {
	GetOrders(ctx context.Context, req domain.GetOrderRequest) ([]domain.Order, error)
	CreateOrder(ctx context.Context, req domain.CreateOrderRequest) (bool, error)
	GetDetailOrder(ctx context.Context, id int) (domain.OrderResponse, error)
	UpdateOrder(ctx context.Context, req domain.UpdateOrderRequest) (bool, error)
}

func NewOrderUsecase(o domain.OrderRepository, p product.ProductRepository, u user.UserRepository) OrderUsecase {
	return &orderUsecase{
		orderRepo:   o,
		productRepo: p,
		userRepo:    u,
	}
}

func (p *orderUsecase) GetOrders(ctx context.Context, req domain.GetOrderRequest) ([]domain.Order, error) {
	orders, err := p.orderRepo.GetAll(req)
	if err != nil {
		log.Errorf("[GetOrders-Usecase] %s", err.Error())
		return []domain.Order{}, err
	}
	return orders, nil
}

func (p *orderUsecase) CreateOrder(ctx context.Context, req domain.CreateOrderRequest) (bool, error) {
	created, err := p.orderRepo.Create(req)
	if err != nil {
		log.Errorf("[CreateOrder-Usecase] %s", err.Error())
		return created, err
	}
	return created, nil
}

func (p *orderUsecase) GetDetailOrder(ctx context.Context, id int) (domain.OrderResponse, error) {
	var wg sync.WaitGroup
	var product product.Product
	var user user.User
	var errProduct, errUser error

	order, err := p.orderRepo.GetById(id)
	if err != nil {
		log.Errorf("[GetDetailOrder-Usecase] %s", err.Error())
		return domain.OrderResponse{}, err
	}

	wg.Add(2)
	go func() {
		product, errProduct = p.productRepo.GetById(order.OrderItemId)
		defer wg.Done()
	}()

	go func() {
		user, errUser = p.userRepo.GetById(order.UserId)
		defer wg.Done()
	}()

	wg.Wait()

	if errProduct != nil {
		log.Errorf("[GetDetailOrder-Usecase] Error Product: %s", errProduct.Error())
		return domain.OrderResponse{}, err
	}

	if errUser != nil {
		log.Errorf("[GetDetailOrder-Usecase] Error User: %s", errUser.Error())
		return domain.OrderResponse{}, err
	}

	orderResponse := domain.OrderResponse{
		ID:           order.ID,
		UserName:     user.FullName,
		ItemName:     product.Name,
		Descriptions: order.Descriptions,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}

	return orderResponse, nil
}

func (p *orderUsecase) UpdateOrder(ctx context.Context, req domain.UpdateOrderRequest) (bool, error) {
	updated, err := p.orderRepo.Update(req)
	if err != nil {
		log.Errorf("[UpdateOrder-Usecase] %s", err.Error())
		return updated, err
	}
	return updated, nil
}
