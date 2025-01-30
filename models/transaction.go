package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Description string
	Entries     []Entry
}

// EntryType represents the type of entry (debit or credit)
type EntryType string

const (
	Debit  EntryType = "DEBIT"
	Credit EntryType = "CREDIT"
)

type Entry struct {
	gorm.Model
	TransactionID uint
	Transaction   Transaction
	AccountID     uint
	Account       Account
	Type          EntryType
	Amount        uint64
}
