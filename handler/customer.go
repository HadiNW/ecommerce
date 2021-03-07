package handler

import (
	"ecommerce-api/domain/customer"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	custService customer.Service
}

func NewHandler(custService customer.Service) *handler {
	return &handler{custService}
}

func (h *handler) RegisterCustomer(c *gin.Context) {
	payload := customer.CustomerRegisterPayload{}
	log.Println("huahahah 1")
	err := c.ShouldBind(&payload)
	if err != nil {
		log.Println("huahahah 2")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.custService.RegisterCustomer(payload)
	if err != nil {
		log.Println("huahahah 3")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("huahahah")

	c.JSON(http.StatusOK, customer.MarshalResponse(created))
}

func (h *handler) ListCustomer(c *gin.Context) {
	data, err := h.custService.ListCustomer()
	if err != nil {
		log.Println("huahahah 3")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("huahahah")

	c.JSON(http.StatusOK, data)
}

func (h *handler) GetCustomer(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	data, err := h.custService.GetCustomerByID(ID)
	if err != nil {
		log.Println("huahahah 3")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("huahahah")

	c.JSON(http.StatusOK, data)
}
