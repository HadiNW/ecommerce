package order

import "time"

type OrderResponse struct {
	ID         int       `json:"id"`
	ProductID  int       `json:"product_id"`
	CartID     int       `json:"cart_id"`
	CustomerID int       `json:"customer_id"`
	Price      int       `json:"price"`
	Qty        int       `json:"qty"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderCreatePayload struct {
	ProductID  int `json:"product_id" binding:"required"`
	Qty        int `json:"qty" binding:"required"`
	CustomerID int
}

func MarshalOrder(o Order) OrderResponse {
	r := OrderResponse{}

	r.ID = o.ID
	r.ProductID = o.ProductID
	r.CartID = o.CartID
	r.CustomerID = o.CustomerID
	r.Price = o.Price
	r.Qty = o.Qty
	r.Status = o.Status
	r.CreatedAt = o.CreatedAt
	r.UpdatedAt = o.UpdatedAt

	return r
}

func MarshalOrders(orders []Order) []OrderResponse {
	r := []OrderResponse{}

	for _, o := range orders {
		data := MarshalOrder(o)
		r = append(r, data)
	}

	return r
}
