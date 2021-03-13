package handler

import (
	"ecommerce-api/domain/product"
	"ecommerce-api/pkg/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetProductByID(c *gin.Context) {

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

func (h *productHandler) ListProduct(c *gin.Context) {
	// ?page=1&limit=10&category=8
	params := product.ProductParam{}

	err := c.ShouldBind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "error params"))
		return
	}

	if params.Search != "" {
		params.Search = "%" + params.Search + "%"
	}

	products, err := h.productService.ListProduct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseBadRequest(err, "error"))
		return
	}

	c.JSON(http.StatusOK, api.ResponseOKPagination(products, nil, "List products success"))
}
