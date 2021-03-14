package order

import (
	"ecommerce-api/domain/product"
	"errors"
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

	p, err := s.productRepo.FindProductByID(order.ProductID)
	if err != nil {
		return Order{}, err
	}

	if p.ID == 0 {
		return Order{}, errors.New("Product not found")
	}

	order.Price = p.Price
	order.Discount = p.Discount
	order.PriceDiscount = p.Price - (p.Price * p.Discount / 100)

	mustAddQty, o, err := s.mustAddQty(order.CustomerID, order.ProductID)
	if err != nil {
		return Order{}, err
	}

	if mustAddQty {
		// Update qty on order data
		order.Qty = o.Qty + payload.Qty
		order.ID = o.ID

		updated, err := s.orderRepo.UpdateQty(order)
		if err != nil {
			return updated, err
		}

		return updated, err
	}

	created, err := s.orderRepo.Create(order)
	if err != nil {
		return created, err
	}

	return created, err
}

func (s *service) mustAddQty(customerID int, productID int) (bool, Order, error) {

	order, err := s.orderRepo.GetOrderByProductAndCustomer(customerID, productID)
	if err != nil {
		return false, order, err
	}

	if order.ID == 0 {
		return false, order, nil
	}

	return true, order, nil
}
