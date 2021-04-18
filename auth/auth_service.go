package auth

import (
	"os"
	"time"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/middleware"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GetAccessToken(userID uint) (string, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) GetAccessToken(userID uint) (string, error) {
	claims := &middleware.JwtCustomClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}
