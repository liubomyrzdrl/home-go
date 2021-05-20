package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string
	Users []*User `gorm:"many2many:users_roles;"`
}
