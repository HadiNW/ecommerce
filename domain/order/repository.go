package order

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(order Order) (Order, error)
	UpdateQty(order Order) (Order, error)
	FindOrderByID(ID int) (Order, error)
	GetCartByCustomer(customerID int) ([]Order, error)
	GetOrderByProductAndCustomer(int, int) (Order, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(order Order) (Order, error) {
	createdOrder := Order{}

	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return createdOrder, err
	}

	// Create ORDER
	query := "INSERT INTO `order` (product_id, cart_id, price, qty, customer_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := tx.Prepare(query)
	if err != nil {
		return createdOrder, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(order.ProductID, 0, order.Price, order.Qty, order.CustomerID)
	if err != nil {
		return createdOrder, err
	}

	err = tx.Commit()
	if err != nil {
		return createdOrder, err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		return createdOrder, err
	}

	createdOrder, err = r.FindOrderByID(int(orderID))
	if err != nil {
		return createdOrder, err
	}

	return createdOrder, nil
}

func (r *repository) UpdateQty(order Order) (Order, error) {
	udpatedOrder := Order{}

	fmt.Printf("%+v\n", order)
	// Begin transaction
	tx := r.db.MustBegin()

	query := "UPDATE `order` SET qty = ? WHERE id = ?"

	tx.MustExec(query, order.Qty, order.ID)

	err := tx.Commit()
	if err != nil {
		return udpatedOrder, err
	}

	udpatedOrder, err = r.FindOrderByID(order.ID)
	if err != nil {
		return udpatedOrder, err
	}
	log.Println(udpatedOrder, "ORRR")

	return udpatedOrder, nil
}

func (r *repository) FindOrderByID(ID int) (Order, error) {
	var order Order
	var orderScan OrderScan
	err := r.db.Get(&orderScan, "SELECT id, product_id, cart_id, customer_id, price, qty, status, created_at, updated_at FROM `order` WHERE id = ?", ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return order, nil
		}
		return order, err
	}
	order.FromScan(orderScan)
	return order, nil
}

func (r *repository) GetCartByCustomer(customerID int) ([]Order, error) {
	query := fmt.Sprintf(`
		select
		id, customer_id, cart_id, product_id, price, qty, status, created_at, updated_at
		from %s
		where status = 'ACTIVE' and customer_id = ?`, "`order`")

	orders := []Order{}
	orderScan := []OrderScan{}

	err := r.db.Select(&orderScan, query, customerID)
	if err != nil {
		return orders, err
	}

	for _, s := range orderScan {
		var o Order
		o.FromScan(s)
		orders = append(orders, o)
	}

	return orders, nil
}

func (r *repository) GetOrderByProductAndCustomer(custID int, productID int) (Order, error) {
	var orderScan OrderScan
	var order Order

	err := r.db.Get(&orderScan, "SELECT * FROM `order` WHERE customer_id = ? AND product_id = ?", custID, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return order, nil
		}
		return order, err
	}

	fmt.Printf("%+v\n", orderScan)

	order.FromScan(orderScan)
	fmt.Printf("%+v\n", order)

	return order, nil
}
