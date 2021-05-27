package service

import (
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/models/repository"
)

type User interface {
	GetAllUsers()(users []models.User, err error)
	CreateUser(user models.User)(err error)
	GetUserById(user models.User, id int64)(err error)
	DeleteUser(user models.User, id int64 )(err error)
}

type Role interface {
	GetRoles() ( roles [] models.Role ,err error)
}

type CreditCard interface {
	GetAllCreditCards() (cards []models.CreditCard, err error)
	CreateCreditCard(card models.CreditCard) (err error)
}

type Service struct {
	User
	Role
	CreditCard
}

func NewService (repo *repository.Repository ) *Service {
	return &Service{
		User: NewUserService(repo.User),
		Role: NewRoleService(repo.Role),
		CreditCard: NewCreditCardService(repo.CreditCard),
	}
}