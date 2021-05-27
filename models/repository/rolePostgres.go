package repository

import (
	"github.com/liubomyrzdrl/home-go/models"
	"gorm.io/gorm"
)

type RolePostgres struct {
	db *gorm.DB
}

// NewRolePostgres TODO Constructor
func NewRolePostgres (db *gorm.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

// GetRoles GetAllRoles GetAllRoles Get all roles  ...  fetch all roles
func (r *RolePostgres) GetRoles() ( roles [] models.Role ,err error) {
	var rs []models.Role
	if err = r.db.Preload("Users").Find(&rs).Error; err != nil {
		return rs, err
	}
	return rs, err
}