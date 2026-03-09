package repository

import "rezafauzan/koda-b6-backend1/internal/models"

type UserRepository struct{
	db *[]models.User
}

var Users []models.User

func NewUserRepository(db *[]models.User) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() []models.User{
	return *r.db
}