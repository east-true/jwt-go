package jwt_test

import (
	"jwt"
	"testing"
)

func TestNewToken(t *testing.T) {
	claims := jwt.NewClaims("auth_token", "test", "boardsvr", "admin")
	token := claims.NewToken()
	if token == "" {
		t.Fail()
	}
}

func TestVerify(t *testing.T) {
	claims := jwt.NewClaims("auth_token", "test", "boardsvr", "admin")
	token := claims.NewToken()
	if token == "" {
		t.Fail()
	}

	newClaims := new(jwt.Claims)
	if !newClaims.Verify(token) {
		t.Fail()
	}
}
