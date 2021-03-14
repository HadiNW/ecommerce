package transaction

import (
	"ecommerce-api/domain/order"
	"errors"
)

type Service interface {
	Checkout([]int, int) (Transaction, error)
}

type service struct {
	transactionRepo Repository
	orderRepo       order.Repository
}

func NewService(transactionRepo Repository, orderRepo order.Repository) Service {
	return &service{transactionRepo, orderRepo}
}

func (s *service) Checkout(orderIDs []int, custID int) (Transaction, error) {

	var t Transaction
	orders, err := s.orderRepo.FindOrderByIDs(orderIDs)
	if err != nil {
		return t, err
	}

	if len(orders) == 0 || len(orders) != len(orderIDs) {
		return t, errors.New("order ids not valid")
	}

	for _, o := range orders {
		if o.CustomerID != custID {
			return t, errors.New("it is not your order")
		}
	}

	t, err = s.transactionRepo.Create(orders, custID)
	if err != nil {
		return t, err
	}

	t.Orders = orders

	return t, nil
}
