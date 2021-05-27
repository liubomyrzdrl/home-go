package service

import (
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/liubomyrzdrl/home-go/models/repository"
)

type RoleService struct {
	repo repository.Role
}

func NewRoleService(repo repository.Role) *RoleService {
	return &RoleService{repo: repo}
}


func (s *RoleService) GetRoles() ( roles [] models.Role ,err error) {
	 return s.repo.GetRoles()
}
