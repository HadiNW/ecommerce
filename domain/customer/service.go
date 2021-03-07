package customer

import "log"

type Service interface {
	RegisterCustomer(customer CustomerRegisterPayload) (Customer, error)
	ListCustomer() ([]Customer, error)
	GetCustomerByID(ID int) (Customer, error)
}

type service struct {
	customerRepo Repository
}

func NewService(customerRepo Repository) Service {
	return &service{customerRepo}
}

func (s *service) RegisterCustomer(customer CustomerRegisterPayload) (Customer, error) {
	created, err := s.customerRepo.Create(customer)
	if err != nil {
		return created, err
	}

	return created, nil
}

func (s *service) ListCustomer() ([]Customer, error) {
	customers, err := s.customerRepo.FindAll()
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (s *service) GetCustomerByID(ID int) (Customer, error) {
	customer, err := s.customerRepo.FindByID(ID)
	log.Println(customer, "CUSTT")
	if err != nil {
		return customer, err
	}

	return customer, nil
}
