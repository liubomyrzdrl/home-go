package models

import (
	"github.com/liubomyrzdrl/home-go/config"
)

// GetAllUsers Get all users ...  fetch all users
func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Preload("CreditCard").Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser Create user ... insert user
func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserById ... fetch user by id
func GetUserById(user *User, id int64) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id int64) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}

// GetAllCreditCards GetAllCreditCards Get all credit cards  ...  fetch all credit cards
func GetAllCreditCards(card *[]CreditCard) (err error) {

	if err = config.DB.Preload("Owner").Find(card).Error; err != nil {
		return err
	}
	return nil
}

// CreateCreditCard Create credit card ... insert card
func CreateCreditCard(card *CreditCard) (err error) {

	if err = config.DB.Create(card).Error; err != nil {
		return err
	}
	return nil
}

// GetRoles GetAllRoles GetAllRoles Get all roles  ...  fetch all roles
func GetRoles(role *[]Role) (err error) {
	if err = config.DB.Preload("Users").Find(role).Error; err != nil {
		return err
	}
	return nil
}

