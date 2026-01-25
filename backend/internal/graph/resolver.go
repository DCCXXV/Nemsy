package graph

import db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	Queries *db.Queries
}
