package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"boosted/livecoding/controllers"
	"boosted/livecoding/models"
	"boosted/livecoding/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Use an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Account{})
	return db
}

func TestAccounts(t *testing.T) {
	// Setup
	db := setupTestDB()
	accountService := services.NewAccountService(db)
	accountController := controllers.NewAccountController(accountService)

	// Create 2 accounts directly in the database
	account1 := models.Account{Name: "Inventory", Type: models.ASSET}
	accountService.CreateAccount(&account1)

	account2 := models.Account{Name: "Accounts Payable", Type: models.LIABILITY}
	accountService.CreateAccount(&account2)

	t.Run("get a single account", func(t *testing.T) {
		// Arrange
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf("/accounts/%d", account1.ID), nil)
		if err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		controllers.RegisterAccountRoutes(mux, accountController)

		// Act
		mux.ServeHTTP(rr, req)

		// Assert
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("get all accounts", func(t *testing.T) {
		// Arrange
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/accounts/", nil)
		if err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		controllers.RegisterAccountRoutes(mux, accountController)

		// Act
		mux.ServeHTTP(rr, req)

		// Assert
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var accounts []models.Account
		json.Unmarshal(rr.Body.Bytes(), &accounts)
		if len(accounts) != 2 {
			t.Errorf("handler returned unexpected body: got %v want %v", accounts, "2")
		}
	})

	t.Run("create an account", func(t *testing.T) {
		// Arrange
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/accounts/", bytes.NewBufferString(`{"name": "Test Account", "type": "ASSET"}`))
		if err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		controllers.RegisterAccountRoutes(mux, accountController)

		// Act
		mux.ServeHTTP(rr, req)

		// Assert

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var account models.Account
		json.Unmarshal(rr.Body.Bytes(), &account)
		if account.Name != "Test Account" {
			t.Errorf("handler returned unexpected body: got %v want %v", account.Name, "Test Account")
		}

		// Validate that the account was created in the database
		db.First(&account)
		if account.Name != "Test Account" {
			t.Errorf("handler returned unexpected body: got %v want %v", account.Name, "Test Account")
		}
	})

	t.Run("update an account", func(t *testing.T) {
		// Arrange
		updatedAccount := models.Account{Name: fmt.Sprintf("Updated %d", rand.Intn(10000)), Type: models.ASSET}
		jsonBodyData, err := json.Marshal(updatedAccount)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("PUT", fmt.Sprintf("/accounts/%d", account2.ID), bytes.NewBuffer(jsonBodyData))
		if err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		controllers.RegisterAccountRoutes(mux, accountController)

		// Act
		mux.ServeHTTP(rr, req)

		// Assert
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var account models.Account
		json.Unmarshal(rr.Body.Bytes(), &account)
		if account.Name != updatedAccount.Name {
			t.Errorf("handler returned unexpected body: got %v want %v", account.Name, updatedAccount.Name)
		}

		db.First(&account)
		if account.Name != updatedAccount.Name {
			t.Errorf("handler returned unexpected body: got %v want %v", account.Name, updatedAccount.Name)
		}
	})
}
