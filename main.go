package main

import (
	"ecommerce/auth"
	"ecommerce/campaign"
	"ecommerce/handler"
	"ecommerce/helper"
	"ecommerce/user"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:adminsaja@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	campaignRepo := campaign.NewRepository(db)

	userService := user.NewService(userRepo)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepo)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	api := router.Group("/api/v1")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "PONG")
	})
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.POST("/users/username-check", userHandler.CheckUsername)
	api.POST("/users/upload-image", authMiddleware(authService, userService), userHandler.UploadImage)

	api.GET("/campaigns", authMiddleware(authService, userService), campaignHandler.ListCampaignByUserID)

	err = router.Run(":9999")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.APIResponseUnAuthorized("Invalid token", errors.New("Invalid token")))
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := authService.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.APIResponseUnAuthorized("Invalid token", err))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.APIResponseUnAuthorized("Invalid token", err))
			return
		}

		userID := int(claims["user_id"].(float64))

		user, err := userService.FindUserByID(userID)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.APIResponseUnAuthorized("User not found", err))
			return
		}

		c.Set("user", user)
	}
}
