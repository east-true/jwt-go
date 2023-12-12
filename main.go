package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte("secret")

type Claim struct {
	Role string
	jwt.RegisteredClaims
}

func NewHmacToken() string {
	claim := new(Claim)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return signedToken
}

func main() {
	token := NewHmacToken()
	if token != "" {
		fmt.Println(token)
	}
}
