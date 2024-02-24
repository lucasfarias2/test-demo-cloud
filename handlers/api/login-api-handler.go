package api

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"log"
	"net/http"
	"packlify-cloud/services"
	"time"
)

type loginRequestData struct {
	Token string `json:"token"`
}

type loginResponseData struct {
	SessionToken string `json:"sessionToken"`
}

func LoginApiHandler(authClient *auth.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

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

		responseData := loginResponseData{
			SessionToken: sessionToken,
		}

		token, err := authClient.VerifySessionCookieAndCheckRevoked(r.Context(), sessionToken)
		if err != nil {
			log.Printf("Failed to verify session cookie: %v", err)
			http.Error(w, "Failed to verify session cookie", http.StatusInternalServerError)
			return
		}

		userAccount, err := services.GetUserAccount(token.UID)
		if err != nil || userAccount == nil {
			_, err := services.CreateAccount(token.UID)
			if err != nil {
				log.Printf("Failed to create account: %v", err)
				http.Error(w, "Failed to create account", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("HX-Redirect", "/dashboard")

		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			log.Printf("Failed to encode response data: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}
	}

}
