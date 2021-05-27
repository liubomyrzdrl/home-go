package service

import (
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/models/repository"
)

type CreditCardService struct {
	repo repository.CreditCard
}

func NewCreditCardService(repo repository.CreditCard) *CreditCardService {
	return &CreditCardService{repo: repo}
}

func (s * CreditCardService) GetAllCreditCards() (cards []models.CreditCard, err error) {
	return  s.repo.GetAllCreditCards()
}


func (s * CreditCardService) CreateCreditCard(card models.CreditCard) (err error) {
	return s.repo.CreateCreditCard(card)
}

//GetAllCreditCards(card []models.CreditCard) (err error)
//CreateCreditCard(card models.CreditCard) (err error)