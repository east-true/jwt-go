package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte("secret")

type Claim struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewHmacToken() string {
	claim := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "test",
			Subject:   "userid",
			Issuer:    "boardsvr",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
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

	claim := new(Claim)
	jwtToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return secret, nil
		}

		return nil, jwt.ErrInvalidType
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if jwtToken.Valid {
		fmt.Println(claim)
	}
}
