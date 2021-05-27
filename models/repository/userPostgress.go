package repository

import (
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

// NewUserPostgres TODO Constructor
func NewUserPostgres (db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// GetAllUsers Get all users ...  fetch all users
func (r *UserPostgres) GetAllUsers() ( users []models.User, err error) {
	var us [] models.User
	if err = r.db.Preload("CreditCard").Find(&us).Error; err != nil {
		return us, err
	}
	return us, nil
}

// CreateUser Create user ... insert user
func (r *UserPostgres) CreateUser(user models.User) (err error) {
	logrus.Println("L", user)
	if err = r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserById ... fetch user by id
func (r *UserPostgres) GetUserById(user models.User, id int64) (err error) {
	if err = r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

//DeleteUser ... Delete user
func (r *UserPostgres) DeleteUser(user models.User, id int64) (err error) {
	r.db.Where("id = ?", id).Delete(&user)
	return nil
}