package studies

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DCCXXV/Nemsy/backend/internal/app"
)

type Handler struct {
	app *app.App
}

func NewHandler(a *app.App) *Handler {
	return &Handler{app: a}
}

func (h *Handler) ListStudies(w http.ResponseWriter, r *http.Request) {
	studies, err := h.app.Queries.ListStudies(r.Context())
	if err != nil {
		log.Printf("Error fetching studies: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studies)
}
