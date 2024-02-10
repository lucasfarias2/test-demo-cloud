package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"log"
	"net/http"
	"os"
)

type User struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func GetUserMiddleware(authClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil || cookie == nil {
				next.ServeHTTP(w, r)
				return
			}

			if os.Getenv("APP_ENV") == "development" && os.Getenv("MOCKS_USER") == "true" {
				user := &User{
					UID:      "1234",
					Email:    "mocked@test.com",
					Name:     "Mocked User",
					Username: "mockedUserName",
				}

				ctx := context.WithValue(r.Context(), "user", user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			token, err := authClient.VerifySessionCookieAndCheckRevoked(r.Context(), cookie.Value)
			if err != nil {
				log.Printf("Failed to verify session cookie: %v", err)
				next.ServeHTTP(w, r)
				return
			}

			userRecord, err := authClient.GetUser(r.Context(), token.UID)
			if err != nil {
				log.Printf("Failed to get user data: %v", err)
				next.ServeHTTP(w, r)
				return
			}

			user := &User{
				UID:      userRecord.UID,
				Email:    userRecord.Email,
				Name:     userRecord.DisplayName,
				Username: userRecord.Email,
			}

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
