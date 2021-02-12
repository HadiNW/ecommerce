package handler

import (
	"ecommerce/helper"
	"ecommerce/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHanlder struct {
	userService user.Service
}

// NewUserHandler ...
func NewUserHandler(userService user.Service) *userHanlder {
	return &userHanlder{userService}
}

func (h *userHanlder) RegisterUser(c *gin.Context) {
	input := user.RegisterInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponseUnprocessable("Malformat JSON", err))
		return
	}
	createdUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Bad request", err))
		return
	}

	c.JSON(http.StatusCreated, helper.APIResponseCreated("user created", user.FormatUser(createdUser)))
}
