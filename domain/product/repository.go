package product

import (
	"database/sql"
	"ecommerce-api/pkg/api"
	"ecommerce-api/pkg/utils"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindProductByID(ID int) (Product, error)
	FindAll(param ProductParam) ([]Product, *api.Pagination, error)
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

func (r *repository) FindAll(param ProductParam) ([]Product, *api.Pagination, error) {
	productScan := []ProductScan{}
	products := []Product{}

	query := "SELECT * FROM product"
	queryPagination := "SELECT COUNT(1) FROM product"
	condition := []string{}

	q, cond := utils.Paginate(utils.Pagination{
		Limit:   param.Limit,
		Offset:  param.Offset,
		OrderBy: param.OrderBy,
		Sort:    param.Sort,
		Search:  param.Search,
	})
	if param.Category != 0 {
		condition = append(condition, "category_id = :category")
	}

	if cond != "" {
		condition = append(condition, cond)
	}

	if len(condition) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(condition, " AND "))
		queryPagination += fmt.Sprintf(" WHERE %s", strings.Join(condition, " AND "))
	}

	query += q
	rows, err := r.db.NamedQuery(query, param)
	if err != nil {
		return products, nil, err
	}

	for rows.Next() {
		scan := ProductScan{}
		err := rows.StructScan(&scan)
		if err != nil {
			return products, nil, err
		}
		productScan = append(productScan, scan)
	}

	for _, p := range productScan {
		var product Product
		product.FromScan(p)
		products = append(products, product)
	}

	var totalData int
	rows, err = r.db.NamedQuery(queryPagination, param)
	if err != nil {
		return products, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&totalData)
		if err != nil {
			return products, nil, err
		}
	}

	pagination := api.Pagination{
		Total:  totalData,
		Limit:  param.Limit,
		Offset: param.Offset,
	}

	return products, &pagination, nil
}
