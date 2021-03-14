package transaction

import (
	"ecommerce-api/domain/order"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(orders []order.Order, custID int) (Transaction, error)
	FindByID(ID int) (Transaction, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(orders []order.Order, custID int) (Transaction, error) {
	query := "INSERT INTO transaction (customer_id, total) VALUES (?, ?)"
	// Begin TX
	tx, err := r.db.Beginx()
	if err != nil {
		return Transaction{}, err
	}

	var total int
	for _, order := range orders {
		total += (order.Price * order.Qty)
	}

	res, err := tx.Exec(query, custID, total)
	if err != nil {
		return Transaction{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return Transaction{}, err
	}

	args := []map[string]interface{}{}
	for _, order := range orders {
		arg := map[string]interface{}{}
		arg["transaction_id"] = lastID
		arg["order_id"] = order.ID
		args = append(args, arg)
	}

	query = "INSERT INTO transaction_detail (transaction_id, order_id) VALUES (:transaction_id, :order_id)"
	_, err = tx.NamedExec(query, args)
	if err != nil {
		return Transaction{}, err
	}

	query = "UPDATE `order` SET status = 'CHECKED_OUT' WHERE id = ?"
	log.Println(query, "qq")
	for _, order := range orders {
		_, err = tx.Exec(query, order.ID)
		if err != nil {
			return Transaction{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return Transaction{}, err
	}

	trx, err := r.FindByID(int(lastID))
	if err != nil {
		return trx, err
	}

	return trx, nil
}

func (r *repository) FindByID(ID int) (Transaction, error) {
	trxScan := TransactionScan{}
	trx := Transaction{}
	err := r.db.Get(&trxScan, "SELECT * FROM transaction WHERE id = ?", ID)
	if err != nil {
		return trx, err
	}

	trx.FromScan(trxScan)

	return trx, nil
}
