package handler

import (
	"ecommerce-api/domain/order"
	"ecommerce-api/domain/product"
	"ecommerce-api/pkg/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService   order.Service
	productService product.Service
}

func NewOrderHandler(orderService order.Service, productService product.Service) *orderHandler {
	return &orderHandler{orderService, productService}
}

func (h *orderHandler) GetCart(c *gin.Context) {
	ID, _ := c.Get("user_id")
	data, err := h.orderService.ListCart(ID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(order.MarshalOrders(data), "success"))
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	id, _ := c.Get("user_id")
	ID := id.(int)

	orderPayload := order.OrderCreatePayload{}

	err := c.ShouldBindJSON(&orderPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	orderPayload.CustomerID = ID

	data, err := h.orderService.CreateOrder(orderPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(data, "success"))
}

func (h *orderHandler) GetProductByID(c *gin.Context) {

	IDstr := c.Param("id")

	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	data, err := h.productService.GetProductByID(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(product.MarshalProduct(data), "success"))
}
