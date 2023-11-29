package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
