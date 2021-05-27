package service

import (
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/models/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func  (s *UserService) GetAllUsers()(users []models.User, err error) {
	return s.repo.GetAllUsers()
}

func  (s *UserService) CreateUser ( user models.User) (err error){
	return s.repo.CreateUser(user)
}

func  (s *UserService)  GetUserById(user models.User, id int64)(err error) {
	return s.repo.GetUserById(user, id)
}

func (s *UserService)  DeleteUser(user models.User, id int64 )( err error) {
	 return s.repo.DeleteUser(user, id)
	//return s.repo.DeleteUser( id)
}