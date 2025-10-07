package studies

import (
	"encoding/json"
	"net/http"

	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
)

type Handler struct {
	db *db.Queries
}

func NewHandler(q *db.Queries) *Handler {
	return &Handler{db: q}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	studies, err := h.db.ListSubjects(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(studies)
}
