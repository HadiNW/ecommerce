package main

import (
	"ecommerce-api/database"
	"ecommerce-api/domain/customer"
	"ecommerce-api/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type IDs struct {
	ID int `db:"id"`
}

func main() {

	db := database.NewDBConnection()

	r := gin.Default()
	api := r.Group("/api/v1")

	custRepo := customer.NewRepository(db)
	custService := customer.NewService(custRepo)
	custHandler := handler.NewHandler(custService)

	api.GET("/customers", custHandler.ListCustomer)
	api.GET("/customers/:id", custHandler.GetCustomer)
	api.POST("/customers/register", custHandler.RegisterCustomer)

	err := r.Run(":9999")
	if err != nil {
		log.Panic(err)
	}
}
