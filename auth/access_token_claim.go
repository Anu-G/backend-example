package auth

import "github.com/golang-jwt/jwt"

type MyClaims struct {
	jwt.StandardClaims
	UserName   string `json:"userName"`
	Email      string
	AccessUuid string `json:"accessUUID"`
}
