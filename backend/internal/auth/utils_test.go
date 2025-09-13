package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestJWTGenerationAndParsing(t *testing.T) {
	secret := []byte("mydevsupersecret")
	tokenStr, err := GenerateJWT("test@ucm.es", secret)
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		t.Fatal("Parse failed:", err)
	}
	if !token.Valid {
		t.Fatal("Token not valid")
	}
}

func TestGenerateStateUnique(t *testing.T) {
	s1 := generateState()
	s2 := generateState()

	if s1 == s2 {
		t.Error("expected states to be different")
	}

	if s1 == "" || s2 == "" {
		t.Error("expected non-empty states")
	}
}

func TestAuthMiddleware_NoCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/protected", nil)
	rr := httptest.NewRecorder()

	mw := &AuthMiddleware{Secret: []byte("testsecret")}
	handler := mw.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}
}
