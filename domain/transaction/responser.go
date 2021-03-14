package transaction

import (
	"ecommerce-api/domain/order"
	"log"
	"time"
)

type TransactionResponse struct {
	ID            int                   `json:"id"`
	CustomerID    int                   `json:"customer_id"`
	Total         int                   `json:"total"`
	PaidAt        *time.Time            `json:"paid_at"`
	PaymentURL    string                `json:"payment_url"`
	PaymentMethod string                `json:"payment_method"`
	Status        string                `json:"status"`
	Orders        []order.OrderResponse `json:"orders"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     *time.Time            `json:"updated_at"`
}
type TransactionPayload struct {
	Orders []int `json:"orders"`
}

func MarshalTransaction(t Transaction) TransactionResponse {
	r := TransactionResponse{}

	orders := order.MarshalOrders(t.Orders)

	r.ID = t.ID
	r.CustomerID = t.CustomerID
	r.Total = t.Total
	r.PaidAt = &t.PaidAt
	r.PaymentURL = t.PaymentURL
	r.PaymentMethod = t.PaymentMethod
	r.Status = t.Status
	r.Orders = orders
	r.CreatedAt = t.CreatedAt
	r.UpdatedAt = &t.UpdatedAt

	log.Println("Marshal tx aja", t.Orders)

	return r
}

func MarshalTransactions(transactions []Transaction) []TransactionResponse {
	r := []TransactionResponse{}

	for _, t := range transactions {
		r = append(r, MarshalTransaction(t))
	}

	log.Println("Marshal trxs")

	return r
}
