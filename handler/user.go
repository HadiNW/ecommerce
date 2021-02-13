package handler

import (
	"ecommerce/auth"
	"ecommerce/helper"
	"ecommerce/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHanlder struct {
	userService user.Service
	authService auth.Service
}

// NewUserHandler ...
func NewUserHandler(userService user.Service, authService auth.Service) *userHanlder {
	return &userHanlder{
		userService,
		authService,
	}
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

	c.JSON(http.StatusCreated, helper.APIResponseCreated("user created", user.FormatUser(createdUser, "")))
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

	token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Bad request", err))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponseOK("Login Success", user.FormatUser(userData, token)))
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

func (h *userHanlder) UploadImage(c *gin.Context) {
	img, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Upload image failed", err))
		return
	}

	decoded, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("User not found", err))
		return
	}

	user := decoded.(user.User)
	userID := user.ID

	path := fmt.Sprintf("images/%d-%s", userID, img.Filename)

	_, err = h.userService.UploadImage(userID, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Upload image failed", err))
		return
	}

	err = c.SaveUploadedFile(img, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("Upload image failed", err))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponseOK("Upload image success", map[string]interface{}{"is_uploaded": true}))
}
