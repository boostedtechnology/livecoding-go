package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Description string
	Entries     []Entry
}
