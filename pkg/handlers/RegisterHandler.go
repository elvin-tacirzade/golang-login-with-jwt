package handlers

import (
	"encoding/json"
	validator "github.com/elvin-tacirzade/golang-validator"
	"golang-login-with-jwt/pkg/config"
	"golang-login-with-jwt/pkg/helpers"
	"golang-login-with-jwt/pkg/models"
	"net/http"
	"time"
)

var db = config.ConnectDB()

func RegisterHandlerFunc(w http.ResponseWriter, r *http.Request) {
	rules := map[string][]string{
		"name":             {"required"},
		"email":            {"required", "email"},
		"password":         {"required", "min:8"},
		"password_confirm": {"required", "same:password"},
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

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	newUser := models.User{}

	res := db.First(&newUser, "email = ?", email)
	if res.RowsAffected == 1 {
		w.WriteHeader(http.StatusBadRequest)
		helpers.CheckError(json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "The email has already been taken.",
		}))
		return
	}

	pass, err := helpers.HashPassword(password)
	helpers.CheckError(err)

	newUser.Name = name
	newUser.Email = email
	newUser.Password = pass
	newUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	result := db.Create(&newUser)
	helpers.CheckError(result.Error)

	data := map[string]string{
		"status":  "success",
		"message": "User successfully register",
	}
	w.WriteHeader(http.StatusOK)
	helpers.CheckError(json.NewEncoder(w).Encode(data))
}
