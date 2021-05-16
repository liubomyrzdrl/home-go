package models

import (
	"github.com/liubomyrzdrl/home-go/config"
)

// GetAllUsers Get all users ...  fetch all users
func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
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
func  GetUserById(user *User, id int64) (err error) {
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