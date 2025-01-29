package services

import (
	"boosted/livecoding/models"

	"gorm.io/gorm"
)

// AccountService provides methods to interact with the Account model
type AccountService struct {
	DB *gorm.DB
}

// NewAccountService creates a new AccountService
func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{DB: db}
}

// CreateAccount creates a new account
func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.DB.Create(account).Error
}

// GetAccount retrieves an account by ID
func (s *AccountService) GetAccount(id uint) (*models.Account, error) {
	var account models.Account
	err := s.DB.First(&account, id).Error
	return &account, err
}

// UpdateAccount updates an existing account
func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.DB.Save(account).Error
}

// ListAccounts retrieves all accounts
func (s *AccountService) ListAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := s.DB.Find(&accounts).Error
	return accounts, err
}
