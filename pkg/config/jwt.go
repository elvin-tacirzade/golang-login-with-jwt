package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SecretKey = "secret_key"
	TokenExp  = time.Now().Add(time.Hour).Unix()
)

type Claims struct {
	jwt.StandardClaims
	UserEmail string
}

func GetUserEmail(tokenStr string) string {
	claims := &Claims{}
	_, _ = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return claims.UserEmail
}

func TokenIsValid(tokeStr string) bool {
	t, err := jwt.Parse(tokeStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return false
	}

	return t.Valid
}
