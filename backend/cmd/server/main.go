package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DCCXXV/Nemsy/backend/internal/auth"
	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type App struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
}

func main() {
	_ = godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	app := &App{
		DB:      pool,
		Queries: queries,
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		log.Fatal("JWT_SECRET not set")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

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
		protected.Get("/api/me", app.meHandler)
	})

	srv := &http.Server{Addr: ":8080", Handler: r}
	go func() {
		log.Println("Server starting on :8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down...")
	srv.Close() // TODO: use Shutdown for prod
}

func (a *App) meHandler(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(auth.CtxUserInfo).(auth.UserInfo)

	if !ok || userInfo.Email == "" {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{"error": "user not found in context"})
		return
	}

	user, err := a.Queries.GetUserByEmail(r.Context(), userInfo.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("User not found, creating new user:", userInfo.Email)

			newUser, err := a.Queries.CreateUser(r.Context(), db.CreateUserParams{
				GoogleSub: userInfo.GoogleSub,
				Email:     userInfo.Email,
				FullName:  stringToPgText(userInfo.FullName),
				Pfp:       stringToPgText(userInfo.Picture),
				Hd:        stringToPgText(userInfo.Hd),
			})

			if err != nil {
				log.Println("Error creating user:", err)
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, map[string]string{"error": "could not create user"})
				return
			}

			render.Status(r, http.StatusCreated)
			render.JSON(w, r, newUser)
			return
		}
		log.Println("Database error:", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "database error"})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func stringToPgText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}
