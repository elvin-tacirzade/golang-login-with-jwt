package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	validator "github.com/elvin-tacirzade/golang-validator"
	"golang-login-with-jwt/pkg/config"
	"golang-login-with-jwt/pkg/helpers"
	"golang-login-with-jwt/pkg/models"
	"net/http"
)

func LoginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	rules := map[string][]string{
		"email":    {"required", "email"},
		"password": {"required"},
	}
	msg := validator.New(r, rules)
	if len(msg) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		helpers.CheckError(json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": msg,
		}))
		return
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := models.User{}

	result := db.First(&user, "email = ?", email)
	if result.RowsAffected == 1 && helpers.CheckPasswordHash(password, user.Password) {

		t := jwt.NewWithClaims(jwt.SigningMethodHS256, config.Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: config.TokenExp,
			},
			UserEmail: email,
		})
		token, err := t.SignedString([]byte(config.SecretKey))
		helpers.CheckError(err)
		w.WriteHeader(http.StatusOK)
		helpers.CheckError(json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"message": "You have successfully logged in",
			"authorization": map[string]interface{}{
				"access_token": token,
				"token_type":   "bearer",
				"expires_in":   "1 hour",
			},
		}))
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	helpers.CheckError(json.NewEncoder(w).Encode(map[string]string{
		"status":  "error",
		"message": "Invalid credentials",
	}))
}
