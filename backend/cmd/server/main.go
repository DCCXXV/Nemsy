package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/DCCXXV/Nemsy/backend/internal/app"
	"github.com/DCCXXV/Nemsy/backend/internal/auth"
	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/DCCXXV/Nemsy/backend/internal/graph"
	"github.com/DCCXXV/Nemsy/backend/internal/studies"
	"github.com/DCCXXV/Nemsy/backend/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	myApp := &app.App{
		Queries: queries,
		DB:      pool,
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		log.Fatal("JWT_SECRET not set")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	store := auth.NewStateStore(5 * time.Minute)

	authHandler := auth.NewHandler(auth.GoogleOAuthConfig(), secret, store, myApp.Queries)
	studiesHandler := studies.NewHandler(myApp)
	usersHandler := users.NewHandler(myApp)

	r.Get("/auth/login", authHandler.LoginHandler)
	r.Get("/auth/callback", authHandler.CallbackHandler)

	mw := &auth.AuthMiddleware{Secret: secret}
	r.Group(func(protected chi.Router) {
		protected.Use(mw.Middleware)
		protected.Get("/api/me", usersHandler.MeHandler)
		protected.Put("/api/me/study", usersHandler.UpdateUserStudy)
		protected.Get("/api/studies", studiesHandler.ListSubjects)

		protected.Handle("/graphql", handler.NewDefaultServer(
			graph.NewExecutableSchema(graph.Config{
				Resolvers: &graph.Resolver{Queries: queries},
			}),
		))
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
