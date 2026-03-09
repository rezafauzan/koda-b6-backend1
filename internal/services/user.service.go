package services

import (
	"rezafauzan/koda-b6-backend1/internal/models"
	"rezafauzan/koda-b6-backend1/internal/repository"
)

type UserService struct{
	repo *repository.UserRepository
}

func NewUserService (repo *repository.UserRepository) *UserService{
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() []models.User{
	return s.repo.GetAll()
}