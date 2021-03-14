package transaction

import (
	"database/sql"
	"ecommerce-api/domain/order"
	"time"
)

type Transaction struct {
	ID            int
	CustomerID    int
	Total         int
	PaidAt        time.Time
	PaymentURL    string
	PaymentMethod string
	Status        string
	Orders        []order.Order
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TransactionDetail struct {
	ID            int
	TransactionID int
	OrderID       int
	Order         *order.Order
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TransactionScan struct {
	ID            sql.NullInt64  `db:"id"`
	CustomerID    sql.NullInt64  `db:"customer_id"`
	Total         sql.NullInt64  `db:"total"`
	PaidAt        sql.NullTime   `db:"paid_at"`
	PaymentURL    sql.NullString `db:"payment_url"`
	PaymentMethod sql.NullString `db:"payment_method"`
	Status        sql.NullString `db:"status"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

type TransactionDetailScan struct {
	ID            sql.NullInt64 `db:"id"`
	TransactionID sql.NullInt64 `db:"transaction_id"`
	OrderID       sql.NullInt64 `db:"order_id"`
	CreatedAt     sql.NullTime  `db:"created_at"`
	UpdatedAt     sql.NullTime  `db:"updated_at"`
}

func (t *Transaction) FromScan(s TransactionScan) {
	t.ID = int(s.CustomerID.Int64)
	t.CustomerID = int(s.CustomerID.Int64)
	t.Total = int(s.Total.Int64)
	t.PaidAt = s.PaidAt.Time
	t.PaymentURL = s.PaymentURL.String
	t.PaymentMethod = s.PaymentMethod.String
	t.Status = s.PaymentURL.String
	t.CreatedAt = s.CreatedAt.Time
	t.UpdatedAt = s.UpdatedAt.Time
}
