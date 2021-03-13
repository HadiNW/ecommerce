package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   bool   `json:"error"`
}

type Response struct {
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type Pagination struct {
	Page   int `json:"page" form:"page"`
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

func ResponseOK(data interface{}, message string) Response {
	return Response{
		Meta: Meta{
			Message: message,
			Code:    200,
			Error:   false,
		},
		Data: data,
	}
}

func ResponseOKPagination(data interface{}, pagination interface{}, message string) Response {
	return Response{
		Meta: Meta{
			Message: message,
			Code:    200,
			Error:   false,
		},
		Data:       data,
		Pagination: &pagination,
	}
}

func ResponseWithCode(data interface{}, message string, code int) Response {
	res := Response{
		Meta: Meta{
			Message: message,
			Code:    code,
			Error:   false,
		},
		Data: data,
	}

	if code > 300 {
		res.Data = errHandler(data.(error))
		res.Meta.Error = true
	}
	return res
}

func ResponseBadRequest(err error, message string) Response {
	return Response{
		Meta: Meta{
			Message: message,
			Code:    400,
			Error:   true,
		},
		Data: errHandler(err),
	}
}

func errHandler(err error) interface{} {
	if err == nil {
		return nil
	}
	var errors []string

	validations, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range validations {
			errors = append(errors, v.Error())
		}
		return gin.H{"errors": errors}
	}

	errors = append(errors, err.Error())
	return gin.H{"errors": errors}

}
