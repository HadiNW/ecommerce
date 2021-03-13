package handler

import (
	"ecommerce-api/domain/auth"
	"ecommerce-api/domain/customer"
	"ecommerce-api/domain/order"
	"ecommerce-api/pkg/api"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	custService  customer.Service
	authService  auth.Service
	orderService order.Service
}

func NewCustomerHandler(custService customer.Service, authService auth.Service, orderService order.Service) *customerHandler {
	return &customerHandler{
		custService,
		authService,
		orderService,
	}
}

func (h *customerHandler) RegisterCustomer(c *gin.Context) {
	payload := customer.CustomerRegisterPayload{}

	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Binding Error"))
		return
	}

	created, err := h.custService.RegisterCustomer(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Create error"))
		return
	}

	c.JSON(http.StatusOK, customer.MarshalResponse(created))
}

func (h *customerHandler) LoginCustomer(c *gin.Context) {
	payload := customer.CustomerLoginPayload{}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Malformat json"))
		return
	}

	cust, err := h.custService.LoginCustomer(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "error"))
		return
	}

	token, err := h.authService.GenerateToken(cust.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "error"))
		return
	}

	response := customer.MarshalResponse(cust)
	response.SetToken(token)

	c.JSON(http.StatusOK, api.ResponseOK(response, "succes login"))
}

func (h *customerHandler) ListCustomer(c *gin.Context) {
	data, err := h.custService.ListCustomer()
	if err != nil {
		log.Println("huahahah 3")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("huahahah")

	c.JSON(http.StatusOK, data)
}

func (h *customerHandler) GetCustomer(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	data, err := h.custService.GetCustomerByID(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "Error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOK(customer.MarshalResponse(data), "success"))
}
