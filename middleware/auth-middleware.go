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

func AuthMiddleware(next http.Handler, authClient *auth.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the session cookie from the request
		cookie, err := r.Cookie("session")
		if err != nil {
			// If there's no cookie, call the next handler without user info
			next.ServeHTTP(w, r)
			return
		}

		// Verify the session cookie
		token, err := authClient.VerifySessionCookieAndCheckRevoked(r.Context(), cookie.Value)
		if err != nil {
			// If the cookie is invalid or revoked, handle according to your needs
			http.Error(w, "Invalid session cookie", http.StatusUnauthorized)
			return
		}

		// Fetch user details from Firebase using the verified token
		userRecord, err := authClient.GetUser(r.Context(), token.UID)
		if err != nil {
			log.Printf("Failed to get user data: %v", err)
			http.Error(w, "Failed to get user data", http.StatusInternalServerError)
			return
		}

		// Create a User object from the userRecord
		user := User{
			UID:      userRecord.UID,
			Email:    userRecord.Email,
			Name:     userRecord.DisplayName,
			Username: userRecord.Email, // Assuming the username is the email for this example
		}

		// Add the User object to the context
		ctx := context.WithValue(r.Context(), "user", user)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
