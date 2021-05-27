package repository

import (
	"github.com/liubomyrzdrl/home-go/config"
	"github.com/liubomyrzdrl/home-go/models"
	"gorm.io/gorm"
)

type CreditCardPostgres struct {
	db *gorm.DB
}

// NewCreditCardPostgres TODO Constructor
func NewCreditCardPostgres (db *gorm.DB) *CreditCardPostgres {
	return &CreditCardPostgres{db: db}
}

// GetAllCreditCards GetAllCreditCards Get all credit cards  ...  fetch all credit cards
func (r * CreditCardPostgres) GetAllCreditCards() (cards []models.CreditCard, err error) {
	var cs []models.CreditCard
	if err = r.db.Preload("Owner").Find(&cs).Error; err != nil {
		return cs, err
	}
	return cs, nil
}

// CreateCreditCard Create credit card ... insert card
func (r * CreditCardPostgres) CreateCreditCard(card models.CreditCard) (err error) {

	if err = config.DB.Create(&card).Error; err != nil {
		return err
	}
	return nil
}