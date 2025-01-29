package models

import (
	"gorm.io/gorm"
)

// EntryType represents the type of entry (debit or credit)
type EntryType string

const (
	Debit  EntryType = "DEBIT"
	Credit EntryType = "CREDIT"
)

type Entry struct {
	gorm.Model
	TransactionID uint
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
	AccountID     uint
	Account       Account   `gorm:"foreignKey:AccountID"`
	Amount        int       `gorm:"not null"`
	Type          EntryType `gorm:"not null"`
}
