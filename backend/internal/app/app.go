package app

import (
	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/DCCXXV/Nemsy/backend/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
	Storage *storage.S3Client
}
