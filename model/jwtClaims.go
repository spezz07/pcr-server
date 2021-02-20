package model

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Uuid    string
	Account string
	jwt.StandardClaims
}
