package main

import (
	"fmt"
	"jwt/auth"
)

func main() {
	claims := auth.NewClaims("auth_token", "test", "boardsvr", "admin")
	token := claims.NewToken()
	if token == "" {
		return
	}

	newClaims := new(auth.Claims)
	if newClaims.Verify(token) {
		fmt.Println(newClaims)
	}
}
