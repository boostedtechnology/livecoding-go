package services

import (
	"boosted/livecoding/models"

	"gorm.io/gorm"
)

// TransactionsService provides methods to interact with the Transaction model
type TransactionsService struct {
	DB *gorm.DB
}

// NewAccountService creates a new AccountService
func NewTransactionsService(db *gorm.DB) *TransactionsService {
	return &TransactionsService{DB: db}
}

// CreateTransaction creates a new transaction
func (s *TransactionsService) CreateTransaction(transaction *models.Transaction) error {
	// TODO: Implement in Task 2
	return nil
}
