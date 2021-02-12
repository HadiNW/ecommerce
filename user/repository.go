package user

import (
	"gorm.io/gorm"
)

// Repository ...
type Repository interface {
	Save(User) (User, error)
	FindByUsername(username string) (User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByUsername(usernname string) (User, error) {
	var user User
	err := r.db.Where("username = ?", usernname).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
