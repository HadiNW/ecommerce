package user

import (
	"gorm.io/gorm"
)

// Repository ...
type Repository interface {
	Save(User) (User, error)
	FindByUsername(username string) (User, error)
	FindByID(int) (User, error)
	Update(int, User) (User, error)
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

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(ID int, user User) (User, error) {
	err := r.db.Where("id = ?", ID).Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
