package product

import "time"

type ProductResponse struct {
	ID               int       `json:"id"`
	CategoryID       int       `json:"category_id"`
	MerchantID       int       `json:"merchant_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Slug             string    `json:"slug"`
	Status           bool      `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func MarshalProduct(p Product) ProductResponse {
	r := ProductResponse{}

	r.ID = p.ID
	r.CategoryID = p.ID
	r.MerchantID = p.MerchantID
	r.Name = p.Name
	r.ShortDescription = p.ShortDescription
	r.Description = p.Description
	r.Stock = p.Stock
	r.Price = p.Price
	r.Slug = p.Slug
	r.Status = p.Status
	r.CreatedAt = p.CreatedAt
	r.UpdatedAt = p.UpdatedAt

	return r
}
