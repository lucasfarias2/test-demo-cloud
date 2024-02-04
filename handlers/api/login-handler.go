package api

import "net/http"

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// read login token from request body
	tokenFromBody := r.Body.token

	// verify token with firebase

	// if token is valid, return a JWT
	// if token is invalid, return an error
	// if token is missing, return an error
	// if token is expired, return an error
	// if token is revoked, return an error
	// if token is disabled, return an error
	// if token is not found, return an error
	// if token is not verified, return an error
	// for rest return generic error
	// return the JWT as JSON and include cookie (session=JWT) in the response
}
