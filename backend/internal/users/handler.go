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

type MeResponse struct {
	ID        int32   `json:"id"`
	Email     string  `json:"email"`
	FullName  *string `json:"fullName,omitempty"`
	Pfp       *string `json:"pfp,omitempty"`
	Hd        *string `json:"hd,omitempty"`
	StudyID   *int32  `json:"studyId,omitempty"`
	StudyName *string `json:"studyName,omitempty"`
}

func (h *Handler) MeHandler(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(auth.CtxUserInfo).(auth.UserInfo)
	if !ok || userInfo.Email == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.app.Queries.GetUserWithStudyByEmail(r.Context(), userInfo.Email)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	resp := MeResponse{
		ID:    user.ID,
		Email: user.Email,
	}
	if user.FullName.Valid {
		resp.FullName = &user.FullName.String
	}
	if user.Pfp.Valid {
		resp.Pfp = &user.Pfp.String
	}
	if user.Hd.Valid {
		resp.Hd = &user.Hd.String
	}
	if user.StudyIDFk.Valid {
		resp.StudyID = &user.StudyIDFk.Int32
		resp.StudyName = &user.StudyName.String
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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
		http.Error(w, "invalid request body", http.StatusBadRequest)
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
		log.Printf("Failed to update study for user %d: %v", userID, err)
		http.Error(w, "failed to update study", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

type SubjectResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Year string `json:"year,omitempty"`
}

func (h *Handler) MySubjects(w http.ResponseWriter, r *http.Request) {
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

	if !user.StudyID.Valid {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]SubjectResponse{})
		return
	}

	subjects, err := h.app.Queries.ListSubjectsByStudy(r.Context(), user.StudyID)
	if err != nil {
		log.Printf("Error fetching subjects: %v", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	resp := make([]SubjectResponse, len(subjects))
	for i, s := range subjects {
		resp[i] = SubjectResponse{
			ID:   s.ID,
			Name: s.Name,
		}
		if s.Year.Valid {
			resp[i].Year = s.Year.String
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
