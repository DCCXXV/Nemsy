package auth

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

type Handler struct {
	OAuthConfig *oauth2.Config
	JWTSecret   []byte
	StateStore  *StateStore
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	email := payload.Claims["email"].(string)

	log.Printf("Handler using secret: %q", h.JWTSecret)
	jwtToken, err := GenerateJWT(email, h.JWTSecret)
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

	http.Redirect(w, r, "http://localhost:5173/", http.StatusSeeOther)
}
