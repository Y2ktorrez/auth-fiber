package src

import (
	"github.com/mutinho/config"
	"github.com/mutinho/src/handler"
	"github.com/mutinho/src/repository"
	"github.com/mutinho/src/service"
)

type Container struct {
	UserRepo    *repository.UserRepository
	UserService *service.UserService
	UserHandler *handler.UserHandler
}

func SetupContainer() *Container {

	//Users
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &Container{
		//Users
		UserRepo:    userRepo,
		UserService: userService,
		UserHandler: userHandler,
	}

}
