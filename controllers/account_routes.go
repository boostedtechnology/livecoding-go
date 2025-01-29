package controllers

import (
	"github.com/gorilla/mux"
)

// RegisterAccountRoutes sets up the account-related routes
func RegisterAccountRoutes(r *mux.Router, accountController *AccountController) {
	// Subrouter for account routes
	accountRouter := r.PathPrefix("/accounts").Subrouter()
	accountRouter.HandleFunc("", accountController.ListAccounts).Methods("GET")
	accountRouter.HandleFunc("", accountController.CreateAccount).Methods("POST")
	accountRouter.HandleFunc("/{id:[0-9]+}", accountController.GetAccount).Methods("GET")
	accountRouter.HandleFunc("/{id:[0-9]+}", accountController.UpdateAccount).Methods("PUT")
}
