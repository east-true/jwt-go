package claims_test

import (
	. "jwt/claims"
	"testing"
)

func TestNewToken(t *testing.T) {
	claims := New("auth_token", "test", "boardsvr", "admin")
	token := claims.NewToken()
	if token == "" {
		t.Fail()
	}
}

func TestVerify(t *testing.T) {
	claims := New("auth_token", "test", "boardsvr", "admin")
	token := claims.NewToken()
	if token == "" {
		t.Fail()
	}

	newClaims := new(Claims)
	if !newClaims.Verify(token) {
		t.Fail()
	}
}
