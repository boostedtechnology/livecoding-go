package controllers

import (
	"net/http"
)

// RegisterAccountRoutes sets up the account-related routes
func RegisterAccountRoutes(r *http.ServeMux, accountController *AccountController) {
	// Subrouter for account routes
	accountRouter := http.NewServeMux()
	accountRouter.HandleFunc("GET /", accountController.ListAccounts)
	accountRouter.HandleFunc("POST /", accountController.CreateAccount)
	accountRouter.HandleFunc("GET /{id}", accountController.GetAccount)
	accountRouter.HandleFunc("PUT /{id}", accountController.UpdateAccount)

	r.Handle("/accounts/", http.StripPrefix("/accounts", accountRouter))
}
