package user

import (
	"log"

	"gorm.io/gorm"
)

// Repository ...
type Repository interface {
	Save(User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Fatal(err.Error())
	}
}
