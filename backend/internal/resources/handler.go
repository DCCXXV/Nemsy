package resources

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/DCCXXV/Nemsy/backend/internal/app"
	"github.com/DCCXXV/Nemsy/backend/internal/auth"
	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Handler struct {
	app       *app.App
	uploadDir string
}

func NewHandler(a *app.App) *Handler {
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}
	os.MkdirAll(uploadDir, 0755)
	return &Handler{app: a, uploadDir: uploadDir}
}

type ResourceResponse struct {
	ID          int32   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	FileURL     string  `json:"fileUrl"`
	FileSize    *int64  `json:"fileSize,omitempty"`
	CreatedAt   string  `json:"createdAt"`
	Owner       *Owner  `json:"owner,omitempty"`
}

type Owner struct {
	ID       int32   `json:"id"`
	FullName *string `json:"fullName,omitempty"`
	Email    string  `json:"email"`
	Pfp      *string `json:"pfp,omitempty"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(auth.CtxUserID).(int32)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// 100MB max for now
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, "file too large or invalid form", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.FormValue("subjectId")
	title := r.FormValue("title")
	description := r.FormValue("description")

	if subjectIDStr == "" || title == "" {
		http.Error(w, "subjectId and title are required", http.StatusBadRequest)
		return
	}

	subjectID, err := strconv.Atoi(subjectIDStr)
	if err != nil {
		http.Error(w, "invalid subjectId", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%d%s", userID, time.Now().UnixNano(), ext)
	filePath := filepath.Join(h.uploadDir, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	written, err := io.Copy(dst, file)
	if err != nil {
		log.Printf("Failed to write file: %v", err)
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}

	fileURL := "/uploads/" + filename

	var desc pgtype.Text
	if description != "" {
		desc = pgtype.Text{String: description, Valid: true}
	}

	resource, err := h.app.Queries.CreateResource(r.Context(), db.CreateResourceParams{
		OwnerID:     userID,
		SubjectID:   int32(subjectID),
		Title:       title,
		Description: desc,
		FileUrl:     fileURL,
		FileSize:    pgtype.Int8{Int64: written, Valid: true},
	})
	if err != nil {
		log.Printf("Failed to create resource: %v", err)
		os.Remove(filePath)
		http.Error(w, "failed to create resource", http.StatusInternalServerError)
		return
	}

	resp := ResourceResponse{
		ID:        resource.ID,
		Title:     resource.Title,
		FileURL:   resource.FileUrl,
		CreatedAt: resource.CreatedAt.Time.Format(time.RFC3339),
	}
	if resource.Description.Valid {
		resp.Description = &resource.Description.String
	}
	if resource.FileSize.Valid {
		resp.FileSize = &resource.FileSize.Int64
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	resource, err := h.app.Queries.GetResourceWithOwner(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "resource not found", http.StatusNotFound)
		return
	}

	resp := ResourceResponse{
		ID:        resource.ID,
		Title:     resource.Title,
		FileURL:   resource.FileUrl,
		CreatedAt: resource.CreatedAt.Time.Format(time.RFC3339),
		Owner: &Owner{
			ID:    resource.OwnerID,
			Email: resource.OwnerEmail,
		},
	}
	if resource.Description.Valid {
		resp.Description = &resource.Description.String
	}
	if resource.FileSize.Valid {
		resp.FileSize = &resource.FileSize.Int64
	}
	if resource.OwnerFullName.Valid {
		resp.Owner.FullName = &resource.OwnerFullName.String
	}
	if resource.OwnerPfp.Valid {
		resp.Owner.Pfp = &resource.OwnerPfp.String
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListBySubject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	resources, err := h.app.Queries.ListResourcesBySubjectWithOwner(r.Context(), int32(id))
	if err != nil {
		log.Printf("Failed to list resources: %v", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	resp := make([]ResourceResponse, len(resources))
	for i, res := range resources {
		resp[i] = ResourceResponse{
			ID:        res.ID,
			Title:     res.Title,
			FileURL:   res.FileUrl,
			CreatedAt: res.CreatedAt.Time.Format(time.RFC3339),
			Owner: &Owner{
				ID:    res.OwnerID,
				Email: res.OwnerEmail,
			},
		}
		if res.Description.Valid {
			resp[i].Description = &res.Description.String
		}
		if res.FileSize.Valid {
			resp[i].FileSize = &res.FileSize.Int64
		}
		if res.OwnerFullName.Valid {
			resp[i].Owner.FullName = &res.OwnerFullName.String
		}
		if res.OwnerPfp.Valid {
			resp[i].Owner.Pfp = &res.OwnerPfp.String
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
