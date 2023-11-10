package schemas

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
