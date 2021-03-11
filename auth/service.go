package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) GenerateToken(userID int) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenStr, err := token.SignedString([]byte("lala"))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *service) ValidateToken(token string) (*jwt.Token, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.GetSigningMethod("HS256") {
			return nil, errors.New("invalid token")
		}
		return []byte("lala"), nil
	})
	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}
