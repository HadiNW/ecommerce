package main

import (
	"ecommerce-api/database"
	"ecommerce-api/domain/auth"
	"ecommerce-api/domain/customer"
	"ecommerce-api/domain/order"
	"ecommerce-api/domain/product"
	"ecommerce-api/domain/transaction"
	"ecommerce-api/handler"
	"ecommerce-api/pkg/api"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.NewDBConnection()

	r := gin.Default()
	api := r.Group("/api/v1")

	custRepo := customer.NewRepository(db)
	orderRepo := order.NewRepository(db)
	productRepo := product.NewRepository(db)
	transactionRepo := transaction.NewRepository(db)

	custService := customer.NewService(custRepo)
	authService := auth.NewService()
	orderService := order.NewService(orderRepo, productRepo)
	productService := product.NewService(productRepo)
	transactionService := transaction.NewService(transactionRepo, orderRepo)

	custHandler := handler.NewCustomerHandler(custService, authService, orderService)
	orderHandler := handler.NewOrderHandler(orderService, productService)
	productHandler := handler.NewProductHandler(productService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	api.GET("/customers", custHandler.ListCustomer)
	api.GET("/customers/:id", custHandler.GetCustomer)
	api.POST("/customers/register", custHandler.RegisterCustomer)
	api.POST("/customers/login", custHandler.LoginCustomer)

	api.GET("/carts", middleware(custService, authService), orderHandler.GetCart)
	api.POST("/carts", middleware(custService, authService), orderHandler.CreateOrder)

	api.GET("/products", productHandler.ListProduct)
	api.GET("/products/:id", productHandler.GetProductByID)

	api.POST("/checkout", middleware(custService, authService), transactionHandler.Checkout)
	api.GET("/transactions", middleware(custService, authService), transactionHandler.GetCustomerTransaction)

	err := r.Run(":9999")
	if err != nil {
		log.Panic(err)
	}
}

func middleware(custService customer.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if !strings.Contains(token, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.ResponseWithCode(errors.New("invalid token 1"), "not authorized", 401))
			return
		}

		token = strings.ReplaceAll(token, "Bearer ", "")

		jwtToken, err := authService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.ResponseWithCode(err, "not authorized", 401))
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		user, err := custService.GetCustomerByID(userID)
		c.Set("user_id", user.ID)
	}
}
