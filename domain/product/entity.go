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
	Price            float64
	Discount         float64
	Slug             string
	Status           bool
	Images           []ProductImage
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ProductImage struct {
	ID        int
	ImageURL  string
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductImageScan struct {
	ID        sql.NullInt64  `db:"id"`
	ImageURL  sql.NullString `db:"image_url"`
	IsPrimary sql.NullInt64  `db:"is_primary"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

type ProductScan struct {
	ID               sql.NullInt64   `db:"id"`
	CategoryID       sql.NullInt64   `db:"category_id"`
	MerchantID       sql.NullInt64   `db:"merchant_id"`
	Name             sql.NullString  `db:"name"`
	ShortDescription sql.NullString  `db:"short_description"`
	Description      sql.NullString  `db:"description"`
	Stock            sql.NullInt64   `db:"stock"`
	Price            sql.NullFloat64 `db:"price"`
	Discount         sql.NullFloat64 `db:"discount"`
	Slug             sql.NullString  `db:"slug"`
	Status           sql.NullInt64   `db:"status"`
	CreatedAt        sql.NullTime    `db:"created_at"`
	UpdatedAt        sql.NullTime    `db:"updated_at"`
}

func (p *Product) FromScan(s ProductScan) {
	p.ID = int(s.ID.Int64)
	p.CategoryID = int(s.CategoryID.Int64)
	p.MerchantID = int(s.MerchantID.Int64)
	p.Name = s.Name.String
	p.ShortDescription = s.ShortDescription.String
	p.Description = s.Description.String
	p.Stock = int(s.Stock.Int64)
	p.Price = s.Price.Float64
	p.Discount = s.Discount.Float64
	p.Slug = s.Slug.String
	p.CreatedAt = s.CreatedAt.Time
	p.UpdatedAt = s.UpdatedAt.Time

	if s.Status.Int64 == 1 {
		p.Status = true
	}
}

func (p *ProductImage) FromScan(s ProductImageScan) {
	p.ID = int(s.ID.Int64)
	p.IsPrimary = false
	p.ImageURL = s.ImageURL.String
	p.CreatedAt = s.CreatedAt.Time
	p.UpdatedAt = s.UpdatedAt.Time
	if s.IsPrimary.Int64 == 1 {
		p.IsPrimary = true
	}
}
