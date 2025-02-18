package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"boosted/livecoding/controllers"
	"boosted/livecoding/models"
	"boosted/livecoding/routes"
	"boosted/livecoding/services"
)

func main() {
	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Account{}, &models.Transaction{}, &models.Entry{})

	// Initialize services and controllers
	accountService := services.NewAccountService(db)
	accountController := controllers.NewAccountController(accountService)

	transactionsService := services.NewTransactionsService(db)
	transactionsController := controllers.NewTransactionsController(transactionsService)

	// Set up router
	r := http.NewServeMux()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Register account routes
	routes.RegisterAccountRoutes(r, accountController)
	routes.RegisterTransactionsRoutes(r, transactionsController)

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
