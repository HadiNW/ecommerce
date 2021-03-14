package transaction

type TransactionResponse struct{}
type TransactionPayload struct {
	Orders []int `json:"orders"`
}
