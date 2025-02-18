package routes

import (
	"net/http"

	"boosted/livecoding/controllers"
)

// RegisterAccountRoutes sets up the account-related routes
func RegisterAccountRoutes(r *http.ServeMux, accountController *controllers.AccountController) {
	// Subrouter for account routes
	accountRouter := http.NewServeMux()
	accountRouter.HandleFunc("GET /", accountController.ListAccounts)
	accountRouter.HandleFunc("POST /", accountController.CreateAccount)
	accountRouter.HandleFunc("GET /{id}", accountController.GetAccount)
	accountRouter.HandleFunc("PUT /{id}", accountController.UpdateAccount)

	r.Handle("/accounts/", http.StripPrefix("/accounts", accountRouter))
}

func RegisterTransactionsRoutes(r *http.ServeMux, transactionsController *controllers.TransactionsController) {
	// Subrouter for transactions routes
	transactionsRouter := http.NewServeMux()
	transactionsRouter.HandleFunc("POST /", transactionsController.CreateTransaction)

	r.Handle("/transactions/", http.StripPrefix("/transactions", transactionsRouter))
}
