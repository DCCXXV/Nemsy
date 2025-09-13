package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DCCXXV/Nemsy/backend/internal/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	secret := []byte(os.Getenv("JWT_SECRET"))
	store := auth.NewStateStore(5 * time.Minute)

	authHandler := &auth.Handler{
		OAuthConfig: auth.GoogleOAuthConfig(),
		JWTSecret:   secret,
		StateStore:  store,
	}

	r.Get("/auth/login", authHandler.LoginHandler)
	r.Get("/auth/callback", authHandler.CallbackHandler)

	mw := &auth.AuthMiddleware{Secret: secret}
	r.Group(func(protected chi.Router) {
		protected.Use(mw.Middleware)
		protected.Get("/api/me", meHandler)
	})

	http.ListenAndServe(":8080", r)
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	email, _ := r.Context().Value(auth.CtxUserEmail).(string)
	if email == "" {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{"error": "user not found"})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"email": email})
}
