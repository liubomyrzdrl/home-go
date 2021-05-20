package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Number    string
	OwnerID   int
	Owner User
	OwnerType string
}
