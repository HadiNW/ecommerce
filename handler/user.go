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

func (h *userHanlder) LoginUser(c *gin.Context) {
	input := user.LoginInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponseUnprocessable("Malformat JSON", err))
		return
	}

	userData, err := h.userService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Bad request", err))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponseOK("Login Success", user.FormatUser(userData)))
}

// CheckUsername ...
func (h *userHanlder) CheckUsername(c *gin.Context) {
	var input user.UsernameCheckInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.APIResponseUnprocessable("Malformat JSON", err))
		return
	}

	ok, err := h.userService.CheckUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Failed checking username", err))
		return
	}
	response := map[string]interface{}{
		"email_available": !ok,
	}
	metaMsg := "Username Available"
	if ok {
		metaMsg = "Username has been taken"
	}
	c.JSON(http.StatusOK, helper.APIResponseOK(metaMsg, response))
}
