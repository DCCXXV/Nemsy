package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DCCXXV/Nemsy/backend/internal/app"
	"github.com/DCCXXV/Nemsy/backend/internal/auth"
	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

type Handler struct {
	app *app.App
}

func NewHandler(a *app.App) *Handler {
	return &Handler{app: a}
}

func (h *Handler) MeHandler(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(auth.CtxUserInfo).(auth.UserInfo)
	if !ok || userInfo.Email == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.app.Queries.GetUserByEmail(r.Context(), userInfo.Email)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateUserStudy(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(auth.CtxUserID)

	if userIDVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userID, ok := userIDVal.(int32)
	if !ok {
		http.Error(w, "invalid user ID type", http.StatusInternalServerError)
		return
	}

	var req struct {
		StudyID int32 `json:"studyId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.StudyID == 0 {
		http.Error(w, "studyId is required", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.app.Queries.UpdateUserStudy(r.Context(), db.UpdateUserStudyParams{
		ID:      userID,
		StudyID: pgtype.Int4{Int32: req.StudyID, Valid: true},
	})

	if err != nil {
		log.Printf("Failed to update study for user %d to study %d: %v", userID, req.StudyID, err)
		http.Error(w, "Failed to update study", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}
