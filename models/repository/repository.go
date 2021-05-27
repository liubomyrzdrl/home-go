package repository

import (
	"github.com/liubomyrzdrl/home-go/models"
	"gorm.io/gorm"
)

type User interface {
	GetAllUsers()([]models.User, error)
	CreateUser(user models.User)(err error)
	GetUserById(user models.User, id int64)(err error)
	DeleteUser(user models.User, id int64 )(err error)
}

type Role interface {
	GetRoles() (users [] models.Role ,err error)
}

type CreditCard interface {
	GetAllCreditCards() (cards []models.CreditCard, err error)
	CreateCreditCard(card models.CreditCard) (err error)

}


type Repository struct {
	User
	Role
	CreditCard
}

func NewRepository (db *gorm.DB ) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		Role: NewRolePostgres(db),
		CreditCard: NewCreditCardPostgres(db),
	}
}