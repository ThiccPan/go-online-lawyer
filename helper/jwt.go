package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/thiccpan/go-online-lawyer/constants"
)

type AuthJWT interface {
	GenerateToken(email string) (string, error)
}

type authJWT struct {
	secret string
}

func NewAuthJWT() *authJWT {
	return &authJWT{
		constants.JWT_SECRET,
	}
}

func (aj *authJWT) GenerateToken(email string) (string, error) {
	claims := &jwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(constants.JWT_SECRET))
}

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}