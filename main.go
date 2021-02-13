package main

import (
	"ecommerce/handler"
	"ecommerce/user"
	"log"
	"net/http"

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
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "PONG")
	})
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.POST("/users/username-check", userHandler.CheckUsername)
	api.POST("/users/upload-image", userHandler.UploadImage)

	err = router.Run(":9999")
	if err != nil {
		log.Fatal(err.Error())
	}
}
