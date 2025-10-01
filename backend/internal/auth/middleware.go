package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const CtxUserInfo ctxKey = "user_info"

type AuthMiddleware struct {
	Secret []byte
}

func (a *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Println("AuthMiddleware: no session cookie")
			http.Error(w, "Unauthorized: no session cookie", http.StatusUnauthorized)
			return
		}
		tokenStr := cookie.Value
		log.Println("AuthMiddleware: got cookie", tokenStr)
		log.Printf("Middleware using secret: %q", a.Secret)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("AuthMiddleware: wrong signing method")
				return nil, jwt.ErrSignatureInvalid
			}
			return a.Secret, nil
		})
		if err != nil {
			log.Println("AuthMiddleware: jwt.Parse error:", err)
		}
		if token == nil {
			log.Println("AuthMiddleware: token is nil")
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("AuthMiddleware: token not valid")
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		log.Println("AuthMiddleware: token is valid")

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Printf("AuthMiddleware: claims = %#v\n", claims)
			userInfo := ExtractUserInfo(claims)
			if userInfo.Email != "" {
				log.Println("AuthMiddleware: user info =", userInfo)
				ctx := context.WithValue(r.Context(), CtxUserInfo, userInfo)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				log.Println("AuthMiddleware: email claim missing")
			}
		} else {
			log.Println("AuthMiddleware: could not cast claims to MapClaims")
		}

		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	})
}
