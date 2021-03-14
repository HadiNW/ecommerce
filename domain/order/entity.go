package order

import (
	"database/sql"
	"time"
)

type Order struct {
	ID            int       `db:"id"`
	ProductID     int       `db:"product_id"`
	CartID        int       `db:"cart_id"`
	CustomerID    int       `db:"customer_id"`
	Price         float64   `db:"price"`
	Discount      float64   `db:"discount"`
	PriceDiscount float64   `db:"price_discount"`
	Qty           int       `db:"qty"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type Cart struct {
	ID         int `db:"id"`
	CustomerID int `db:"customer_id"`
	Order      []Order
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type OrderScan struct {
	ID            sql.NullInt64   `db:"id"`
	ProductID     sql.NullInt64   `db:"product_id"`
	CustomerID    sql.NullInt64   `db:"customer_id"`
	CartID        sql.NullInt64   `db:"cart_id"`
	Price         sql.NullFloat64 `db:"price"`
	PriceDiscount sql.NullFloat64 `db:"price_discount"`
	Discount      sql.NullFloat64 `db:"discount"`
	Qty           sql.NullInt64   `db:"qty"`
	Status        sql.NullString  `db:"status"`
	CreatedAt     sql.NullTime    `db:"created_at"`
	UpdatedAt     sql.NullTime    `db:"updated_at"`
}

type CartScan struct {
	ID         sql.NullInt64 `db:"id"`
	CustomerID sql.NullInt64 `db:"customer_id"`
	CreatedAt  sql.NullTime  `db:"created_at"`
	UpdatedAt  sql.NullTime  `db:"updated_at"`
}

func (o *Order) FromScan(s OrderScan) {
	o.ID = int(s.ID.Int64)
	o.ProductID = int(s.ProductID.Int64)
	o.CartID = int(s.CartID.Int64)
	o.CustomerID = int(s.CustomerID.Int64)
	o.Price = s.Price.Float64
	o.PriceDiscount = s.PriceDiscount.Float64
	o.Qty = int(s.Qty.Int64)
	o.Discount = s.Discount.Float64
	o.Status = s.Status.String
	o.CreatedAt = s.CreatedAt.Time
	o.UpdatedAt = s.UpdatedAt.Time
}

func (c *Cart) FromScan(s CartScan) {
	c.ID = int(s.ID.Int64)
	c.CustomerID = int(s.CustomerID.Int64)
	c.CreatedAt = s.CreatedAt.Time
	c.UpdatedAt = s.UpdatedAt.Time
}
