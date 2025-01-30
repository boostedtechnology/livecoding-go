package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"boosted/livecoding/models"
	"boosted/livecoding/services"
)

// AccountController handles HTTP requests for accounts
type AccountController struct {
	Service *services.AccountService
}

// NewAccountController creates a new AccountController
func NewAccountController(service *services.AccountService) *AccountController {
	return &AccountController{Service: service}
}

// CreateAccount handles POST /accounts
func (c *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.Service.CreateAccount(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// GetAccount handles GET /accounts/{id}
func (c *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error parsing account ID: %v", err)
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := c.Service.GetAccount(uint(id))
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(account)
}

// UpdateAccount handles PUT /accounts/{id}
func (c *AccountController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error parsing account ID: %v", err)
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	account.ID = uint(id)
	if err := c.Service.UpdateAccount(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(account)
}

// ListAccounts handles GET /accounts
func (c *AccountController) ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := c.Service.ListAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(accounts)
}
