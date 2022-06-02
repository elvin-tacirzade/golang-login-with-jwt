package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang-login-with-jwt/pkg/handlers"
	"golang-login-with-jwt/pkg/helpers"
	"golang-login-with-jwt/pkg/middleware"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	//register
	api.HandleFunc("/register", handlers.RegisterHandlerFunc).Methods("POST")
	//login
	api.HandleFunc("/login", handlers.LoginHandlerFunc).Methods("POST")
	//users
	api.HandleFunc("/users", middleware.IsValid(handlers.UserHandlerFunc)).Methods("GET")

	fmt.Println("Server starting...")
	helpers.CheckError(http.ListenAndServe(":8080", r))
}
