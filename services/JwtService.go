package services

import (
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type JwtService struct {
}

func (jwtService *JwtService) CreateToken(issuer uint, secretKey string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(issuer)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		return "error", err
	}

	return token, nil
}

func (jwtService *JwtService) GetToken(jwtToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (jwtService *JwtService) GetClaims(token *jwt.Token) *jwt.StandardClaims {
	return token.Claims.(*jwt.StandardClaims)
}
