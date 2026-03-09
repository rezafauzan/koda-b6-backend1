package repository

import "rezafauzan/koda-b6-backend1/internal/models"

type UserRepository struct{
	db *[]models.User
}

var Users []models.User

func NewUserRepository() *UserRepository{
	return &UserRepository{
		db: &Users,
	}
}