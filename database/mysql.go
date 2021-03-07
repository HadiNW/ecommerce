package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDBConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:adminsaja@tcp(127.0.0.1:3306)/ecommerce_store?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Panic(err.Error())
	}

	return db
}
