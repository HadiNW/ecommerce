package handler

import (
	"ecommerce-api/domain/transaction"
	"ecommerce-api/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) Checkout(c *gin.Context) {
	payload := transaction.TransactionPayload{}

	id, _ := c.Get("user_id")
	ID := id.(int)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	t, err := h.transactionService.Checkout(payload.Orders, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(transaction.MarshalTransaction(t), "success checkout"))
}

func (h *transactionHandler) GetCustomerTransaction(c *gin.Context) {
	id, _ := c.Get("user_id")
	ID := id.(int)

	transactions, err := h.transactionService.GetCustomerTransactions(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(transaction.MarshalTransactions(transactions), "success"))
}
