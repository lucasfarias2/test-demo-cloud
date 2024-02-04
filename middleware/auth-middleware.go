package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"log"
	"net/http"
)

type User struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func AuthMiddleware(next http.Handler, authClient *auth.Client, forceAuth bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		if forceAuth && cookie == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		token, err := authClient.VerifySessionCookieAndCheckRevoked(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "Invalid session cookie", http.StatusUnauthorized)
			return
		}

		userRecord, err := authClient.GetUser(r.Context(), token.UID)
		if err != nil {
			log.Printf("Failed to get user data: %v", err)
			http.Error(w, "Failed to get user data", http.StatusInternalServerError)
			return
		}

		user := User{
			UID:      userRecord.UID,
			Email:    userRecord.Email,
			Name:     userRecord.DisplayName,
			Username: userRecord.Email,
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
