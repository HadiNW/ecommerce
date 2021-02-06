package user

import "golang.org/x/crypto/bcrypt"

// Service ...
type Service interface {
	RegisterUser(input RegisterInput) (User, error)
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
