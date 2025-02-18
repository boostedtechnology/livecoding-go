package controllers

import (
	"net/http"

	"boosted/livecoding/services"
)

// TransactionsController handles HTTP requests
type TransactionsController struct {
	Service *services.TransactionsService
}

// NewTransactionsController creates a new TransactionsController
func NewTransactionsController(service *services.TransactionsService) *TransactionsController {
	return &TransactionsController{Service: service}
}

// CreateTransaction handles POST /transactions
func (c *TransactionsController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement in Task 2
}
