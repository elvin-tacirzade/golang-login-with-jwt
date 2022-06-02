package middleware

import (
	"encoding/json"
	"golang-login-with-jwt/pkg/config"
	"golang-login-with-jwt/pkg/helpers"
	"net/http"
)

func IsValid(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.TokenIsValid(r.Header.Get("Authorization")) {
			h(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			helpers.CheckError(json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Unauthorized access",
			}))
		}
	}
}
