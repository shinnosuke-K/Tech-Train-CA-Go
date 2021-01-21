package auth

import "github.com/dgrijalva/jwt-go"

func CreateJwt(claims jwt.MapClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}
