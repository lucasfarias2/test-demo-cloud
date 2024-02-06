package api

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"log"
	"net/http"
	"time"
)

type loginRequestData struct {
	Token string `json:"token"`
}

type loginResponseData struct {
	SessionToken string `json:"sessionToken"`
}

func HandleLogin(authClient *auth.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		if r.Method != "POST" {
			http.Error(w, "HTTP Method not accepted", http.StatusMethodNotAllowed)
			return
		}

		var body loginRequestData

		body.Token = r.FormValue("token")

		expiresIn := time.Hour * 24 * 14

		sessionToken, err := authClient.SessionCookie(ctx, body.Token, expiresIn)
		if err != nil {
			log.Printf("Failed to create a session cookie: %v", err)
			http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionToken,
			Path:     "/",
			Expires:  time.Now().Add(expiresIn),
			HttpOnly: true,
			Secure:   true,
		})

		// Prepare the response data with the session token
		responseData := loginResponseData{
			SessionToken: sessionToken,
		}

		// Set response header to application/json
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("HX-Redirect", "/dashboard")

		// Encode and write the response data as JSON
		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			log.Printf("Failed to encode response data: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}
	}

}
