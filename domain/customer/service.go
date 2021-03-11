package customer

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterCustomer(customer CustomerRegisterPayload) (Customer, error)
	ListCustomer() ([]Customer, error)
	GetCustomerByID(ID int) (Customer, error)
	GetCustomerByEmail(email string) (Customer, error)
	LoginCustomer(payload CustomerLoginPayload) (Customer, error)
}

type service struct {
	customerRepo Repository
}

func NewService(customerRepo Repository) Service {
	return &service{customerRepo}
}

func (s *service) RegisterCustomer(customer CustomerRegisterPayload) (Customer, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		return Customer{}, err
	}

	customer.Password = string(hashed)
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
	if err != nil {
		return customer, err
	}

	if customer.ID == 0 {
		return customer, errors.New("User not found")
	}

	return customer, nil
}

func (s *service) GetCustomerByEmail(email string) (Customer, error) {
	customer, err := s.customerRepo.FindByEmail(email)
	if err != nil {
		return customer, err
	}

	if customer.ID == 0 {
		return customer, errors.New("User not found")
	}

	return customer, nil
}

func (s *service) LoginCustomer(c CustomerLoginPayload) (Customer, error) {
	cust, err := s.GetCustomerByEmail(c.Email)
	if err != nil {
		return cust, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(cust.Password), []byte(c.Password))
	if err != nil {
		return cust, err
	}

	return cust, nil
}
