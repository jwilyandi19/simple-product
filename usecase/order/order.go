package order

import (
	"context"

	domain "github.com/jwilyandi19/simple-product/domain/order"
	log "github.com/sirupsen/logrus"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
}

type OrderUsecase interface {
	GetOrders(ctx context.Context, req domain.GetOrderRequest) ([]domain.Order, error)
	CreateOrder(ctx context.Context, req domain.CreateOrderRequest) (bool, error)
	GetDetailOrder(ctx context.Context, id int) (domain.Order, error)
	UpdateOrder(ctx context.Context, req domain.UpdateOrderRequest) (bool, error)
}

func NewOrderUsecase(p domain.OrderRepository) OrderUsecase {
	return &orderUsecase{
		orderRepo: p,
	}
}

func (p *orderUsecase) GetOrders(ctx context.Context, req domain.GetOrderRequest) ([]domain.Order, error) {
	orders, err := p.orderRepo.GetAll()
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

func (p *orderUsecase) GetDetailOrder(ctx context.Context, id int) (domain.Order, error) {
	order, err := p.orderRepo.GetById(id)
	if err != nil {
		log.Errorf("[GetDetailOrder-Usecase] %s", err.Error())
		return domain.Order{}, err
	}
	return order, nil
}

func (p *orderUsecase) UpdateOrder(ctx context.Context, req domain.UpdateOrderRequest) (bool, error) {
	updated, err := p.orderRepo.Update(req)
	if err != nil {
		log.Errorf("[UpdateOrder-Usecase] %s", err.Error())
		return updated, err
	}
	return updated, nil
}
