package handler

import (
	"ecommerce-api/domain/transaction"
	"ecommerce-api/pkg/api"
	"log"
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

	log.Println(payload, "PAYLOAD")

	t, err := h.transactionService.Checkout(payload.Orders, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(t, "success checkout"))

}
