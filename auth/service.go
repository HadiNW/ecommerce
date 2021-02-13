package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("lala"))
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
