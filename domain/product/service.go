package product

import "errors"

type Service interface {
	GetProductByID(int) (Product, error)
	ListProduct() ([]Product, error)
}

type service struct {
	productRepo Repository
}

func NewService(productRepo Repository) Service {
	return &service{productRepo}
}

func (s *service) GetProductByID(ID int) (Product, error) {
	product, err := s.productRepo.FindProductByID(ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("product not found")
	}
	return product, nil
}
func (s *service) ListProduct() ([]Product, error) {
	return []Product{}, nil
}
