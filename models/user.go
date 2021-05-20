package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Username string `json:"username"`
	CreditCard []CreditCard `gorm:"foreignKey: OwnerID;"`
	Role []*Role `gorm:"many2many:users_roles;"`
}


