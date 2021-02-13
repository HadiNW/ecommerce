package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service ...
type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	CheckUsername(username string) (bool, error)
	UploadImage(userID int, location string) (User, error)
	FindUserByID(userID int) (User, error)
}

type service struct {
	repository Repository
}

// NewService ...
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}
	user := User{
		Username: input.Username,
		Password: string(hash),
		FullName: input.FullName,
	}
	createdUser, err := s.repository.Save(user)
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {
	user, err := s.repository.FindByUsername(input.Username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (s *service) CheckUsername(username string) (bool, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return false, nil
	}
	return true, nil
}

func (s *service) UploadImage(userID int, loc string) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	user.Image = loc

	user, err = s.repository.Update(userID, user)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *service) FindUserByID(userID int) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}
