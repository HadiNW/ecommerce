package product

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindProductByID(ID int) (Product, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}

func (r *repository) FindProductByID(ID int) (Product, error) {

	productScan := ProductScan{}
	product := Product{}

	err := r.db.Get(&productScan, "SELECT * FROM product WHERE id = ?", ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, nil
		}
		return product, err
	}

	product.FromScan(productScan)

	return product, nil
}
