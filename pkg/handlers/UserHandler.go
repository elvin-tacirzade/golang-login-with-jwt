package handlers

import (
	"encoding/json"
	"golang-login-with-jwt/pkg/helpers"
	"golang-login-with-jwt/pkg/models"
	"net/http"
)

func UserHandlerFunc(w http.ResponseWriter, _ *http.Request) {
	var users []models.User
	db.Find(&users)
	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		helpers.CheckError(json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "There were no results",
		}))
		return
	}

	w.WriteHeader(http.StatusOK)

	var usersVisible []models.UserVisible
	for _, value := range users {
		usersVisible = append(usersVisible, value.UserToUserVisible())
	}

	helpers.CheckError(json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"users":  usersVisible,
	}))
}
