package order

import (
	"ecommerce-api/domain/product"
)

type Service interface {
	ListCart(customerID int) ([]Order, error)
	CreateOrder(order OrderCreatePayload) (Order, error)
}

type service struct {
	orderRepo   Repository
	productRepo product.Repository
}

func NewService(orderRepo Repository, productRepo product.Repository) Service {
	return &service{orderRepo, productRepo}
}

func (s *service) ListCart(customerID int) ([]Order, error) {
	cart, err := s.orderRepo.GetCartByCustomer(customerID)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (s *service) CreateOrder(payload OrderCreatePayload) (Order, error) {
	order := Order{}

	order.ProductID = payload.ProductID
	order.Qty = payload.Qty
	order.CustomerID = payload.CustomerID

	created, err := s.orderRepo.Create(order)
	if err != nil {
		return created, err
	}

	return created, err
}
