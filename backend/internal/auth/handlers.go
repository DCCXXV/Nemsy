package auth

import (
	"context"
	"log"
	"net/http"

	db "github.com/DCCXXV/Nemsy/backend/internal/db/generated"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

type Handler struct {
	OAuthConfig *oauth2.Config
	JWTSecret   []byte
	StateStore  *StateStore
	Queries     *db.Queries
}

type UserInfo struct {
	GoogleSub string
	Email     string
	FullName  string
	Picture   string
	Hd        string
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	state := generateState()

	h.StateStore.Put(state)

	url := h.OAuthConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("prompt", "select_account"),
	)
	log.Println("Redirecting to:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	returnedState := r.URL.Query().Get("state")
	log.Println("Callback state param:", returnedState)

	if !h.StateStore.Check(returnedState) {
		http.Error(w, "Invalid OAuth state", http.StatusUnauthorized)
		return
	}

	code := r.URL.Query().Get("code")

	token, err := h.OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Exchange failed:", err)
		http.Error(w, "Failed exchange", http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "ID Token missing in OAuth2 response", http.StatusInternalServerError)
		return
	}

	payload, err := idtoken.Validate(context.Background(), rawIDToken, h.OAuthConfig.ClientID)
	if err != nil {
		http.Error(w, "Invalid ID Token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	userInfo := ExtractUserInfo(payload.Claims)
	newUser := false
	if _, err = h.Queries.GetUserByEmail(r.Context(), userInfo.Email); err == pgx.ErrNoRows {
		log.Println("User not found, creating new user:", userInfo.Email)
		_, err = h.Queries.CreateUser(r.Context(), db.CreateUserParams{
			GoogleSub: userInfo.GoogleSub,
			StudyID:   pgtype.Int4{Valid: false},
			Email:     userInfo.Email,
			FullName:  stringToPgText(userInfo.FullName),
			Pfp:       stringToPgText(userInfo.Picture),
			Hd:        stringToPgText(userInfo.Hd),
		})
		newUser = true
	}

	if err != nil {
		log.Println("Database error:", err)
		http.Error(w, "Could not save user", http.StatusInternalServerError)
		return
	}

	log.Printf("Handler using secret: %q", h.JWTSecret)
	jwtToken, err := GenerateJWT(userInfo, h.JWTSecret)
	if err != nil {
		http.Error(w, "Could not create session token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                // TODO: true in prod
		SameSite: http.SameSiteLaxMode, // TODO: None in prod
	})
	if newUser == true {
		http.Redirect(w, r, "http://localhost:5173/auth", http.StatusSeeOther)
	} else {
		// TODO: remember later that if a user closes mid onboarding he isnt "new" but wont be able to choose studies
		http.Redirect(w, r, "http://localhost:5173/", http.StatusSeeOther)
	}
}
