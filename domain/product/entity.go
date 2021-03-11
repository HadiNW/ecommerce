package product

import (
	"database/sql"
	"time"
)

type Product struct {
	ID               int
	CategoryID       int
	MerchantID       int
	Name             string
	ShortDescription string
	Description      string
	Stock            int
	Price            int
	Slug             string
	Status           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ProductScan struct {
	ID               sql.NullInt64  `db:"id"`
	CategoryID       sql.NullInt64  `db:"category_id"`
	MerchantID       sql.NullInt64  `db:"merchant_id"`
	Name             sql.NullString `db:"name"`
	ShortDescription sql.NullString `db:"short_description"`
	Description      sql.NullString `db:"description"`
	Stock            sql.NullInt64  `db:"stock"`
	Price            sql.NullInt64  `db:"price"`
	Slug             sql.NullString `db:"slug"`
	Status           sql.NullInt64  `db:"status"`
	CreatedAt        sql.NullTime   `db:"created_at"`
	UpdatedAt        sql.NullTime   `db:"updated_at"`
}

func (p *Product) FromScan(s ProductScan) {
	p.ID = int(s.ID.Int64)
	p.CategoryID = int(s.CategoryID.Int64)
	p.MerchantID = int(s.MerchantID.Int64)
	p.Name = s.Name.String
	p.ShortDescription = s.ShortDescription.String
	p.Description = s.Description.String
	p.Stock = int(s.Stock.Int64)
	p.Price = int(s.Price.Int64)
	p.Slug = s.Slug.String
	p.CreatedAt = s.CreatedAt.Time
	p.UpdatedAt = s.UpdatedAt.Time

	if s.Status.Int64 == 1 {
		p.Status = true
	}
}
