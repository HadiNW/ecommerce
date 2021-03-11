package customer

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(customer CustomerRegisterPayload) (Customer, error)
	FindByID(int) (Customer, error)
	FindAll() ([]Customer, error)
	FindByEmail(string) (Customer, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (Customer, error) {
	customerScan := CustomerScan{}
	customer := Customer{}

	query := "SELECT id, full_name, email, username, avatar, status, password, created_at, updated_at FROM customer WHERE id = ?"

	err := r.db.Get(&customerScan, query, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Customer{}, nil
		}
		return customer, err
	}

	customer.FromScan(customerScan)

	return customer, nil
}

func (r *repository) FindByEmail(email string) (Customer, error) {
	customerScan := CustomerScan{}
	customer := Customer{}

	query := "SELECT id, full_name, email, username, avatar, status, password, created_at, updated_at FROM customer WHERE email = ?"

	err := r.db.Get(&customerScan, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return Customer{}, nil
		}
		return customer, err
	}

	customer.FromScan(customerScan)

	return customer, nil
}

func (r *repository) FindAll() ([]Customer, error) {
	customerScan := []CustomerScan{}
	customers := []Customer{}

	query := "SELECT id, full_name, email, username, avatar, status, created_at, updated_at FROM customer"

	err := r.db.Select(&customerScan, query)
	if err != nil {
		return customers, err
	}

	for _, s := range customerScan {
		c := Customer{}
		c.FromScan(s)
		customers = append(customers, c)
	}

	return customers, nil
}

func (r *repository) Create(customer CustomerRegisterPayload) (Customer, error) {
	query := "INSERT INTO customer (full_name, email, password, username) VALUES (?, ?, ?, ?)"

	res, err := r.db.Exec(query, customer.FullName, customer.Email, customer.Password, customer.Email)
	if err != nil {
		return Customer{}, err
	}

	lastID, err := res.LastInsertId()

	data, err := r.FindByID(int(lastID))
	if err != nil {
		return Customer{}, err
	}

	return data, nil
}
