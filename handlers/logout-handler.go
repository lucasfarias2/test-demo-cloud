package handlers

import (
	"net/http"
	"time"
)

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:    "session",
			Value:   "",
			Expires: time.Unix(0, 0),
			Path:    "/",
			MaxAge:  -1,
		}

		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
