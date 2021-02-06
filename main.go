package main

import (
	"ecommerce/user"
	"log"

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
	createdUser, err := userService.RegisterUser(user.RegisterInput{
		Username: "user3",
		Password: "qwe",
		FullName: "User 3",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(createdUser)
}
