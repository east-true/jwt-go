package jwt

import (
	"fmt"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte("secret")

type Claims struct {
	Role string `json:"role"`
	gojwt.RegisteredClaims
}

func NewClaims(id, sub, iss, role string) *Claims {
	return &Claims{
		Role: role,
		RegisteredClaims: gojwt.RegisteredClaims{
			ID:        id,
			Subject:   sub,
			Issuer:    iss,
			IssuedAt:  gojwt.NewNumericDate(time.Now()),
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
}

func (c *Claims) NewToken() string {
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return signedToken
}

func (c *Claims) Verify(token string) bool {
	jwtToken, err := gojwt.ParseWithClaims(token, c, func(t *gojwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*gojwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		exp, err := t.Claims.GetExpirationTime()
		if err != nil {
			return nil, gojwt.ErrTokenInvalidClaims
		}

		if exp.After(time.Now()) {
			return nil, gojwt.ErrTokenExpired
		}

		return nil, gojwt.ErrInvalidKeyType
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	claims := jwtToken.Claims.(*Claims)
	c.Role = claims.Role
	c.RegisteredClaims = claims.RegisteredClaims
	return jwtToken.Valid
}
