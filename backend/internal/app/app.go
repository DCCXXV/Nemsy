package app

import (
	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
}
