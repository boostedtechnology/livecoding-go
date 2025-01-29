package models

import "gorm.io/gorm"

// AccountType represents the type of account
type AccountType string

const (
	ASSET     AccountType = "ASSET"
	LIABILITY AccountType = "LIABILITY"
	EQUITY    AccountType = "EQUITY"
)

// Account represents a financial account in a ledger
type Account struct {
	gorm.Model
	Name string      `gorm:"not null"`
	Type AccountType `gorm:"not null"`
}
