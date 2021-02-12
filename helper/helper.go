package helper

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Meta ...
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// Response ...
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// APIResponseWithCode ...
func APIResponseWithCode(code int, message string, status string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}

// APIResponseCreated ...
func APIResponseCreated(message string, data interface{}) Response {
	return APIResponseWithCode(http.StatusCreated, message, "Success", data)
}

// APIResponseOK ...
func APIResponseOK(message string, data interface{}) Response {
	return APIResponseWithCode(http.StatusOK, message, "Success", data)
}

// APIResponseBadRequest ...
func APIResponseBadRequest(message string, err error) Response {
	return APIResponseWithCode(http.StatusBadRequest, message, "Bad Request", errFormatter(err))
}

// APIResponseUnprocessable ...
func APIResponseUnprocessable(message string, err error) Response {
	return APIResponseWithCode(http.StatusUnprocessableEntity, message, "Unprocessable Entity", errFormatter(err))
}

func errFormatter(err error) interface{} {
	var errors []string
	type H map[string]interface{}

	_, ok := err.(validator.ValidationErrors)
	if !ok {
		errors = append(errors, err.Error())
		return H{"errors": errors}
	}
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return H{"errors": errors}
}
