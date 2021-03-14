package product

import "time"

type ProductResponse struct {
	ID               int            `json:"id"`
	CategoryID       int            `json:"category_id"`
	MerchantID       int            `json:"merchant_id"`
	Name             string         `json:"name"`
	ShortDescription string         `json:"short_description"`
	Description      string         `json:"description"`
	Stock            int            `json:"stock"`
	Price            float64        `json:"price"`
	Discount         float64        `json:"discount"`
	Slug             string         `json:"slug"`
	Status           bool           `json:"status"`
	Images           []ProductImage `json:"images"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

type PaginationPayload struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ProductParam struct {
	Page     int    `json:"page" form:"page"`
	Limit    int    `json:"limit" form:"limit"`
	Offset   int    `json:"offset" form:"offset"`
	Search   string `json:"search" form:"search"`
	OrderBy  string `json:"order_by" form:"order"`
	Sort     string `json:"sort" form:"sort"`
	Category int    `json:"category" form:"category"`
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
	r.Discount = p.Discount
	r.Slug = p.Slug
	r.Status = p.Status
	r.Images = p.Images
	r.CreatedAt = p.CreatedAt
	r.UpdatedAt = p.UpdatedAt

	return r
}

func MarshalProducts(products []Product) []ProductResponse {
	r := []ProductResponse{}

	for _, product := range products {
		r = append(r, MarshalProduct(product))
	}

	return r
}
