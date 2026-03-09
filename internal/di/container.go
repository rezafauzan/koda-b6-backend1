package container

import (
	"rezafauzan/koda-b6-backend1/internal/handlers"
	"rezafauzan/koda-b6-backend1/internal/models"
	"rezafauzan/koda-b6-backend1/internal/repository"
	"rezafauzan/koda-b6-backend1/internal/services"
)

type Container struct {
	users          *[]models.User
	userRepository *repository.UserRepository
	userService    *services.UserService
	userHandlers   *handlers.UserHandler
}

func NewContainer() *Container {
	var Users []models.User
	container := Container{
		users: &Users,
	}
	return &container
}

func (c *Container) initDependencies() {
c.userRepository = repository.NewUserRepository(c.users)
c.userService = services.NewUserService(c.userRepository)
c.userHandlers = handlers.NewUserHandler(c.userService)
}
