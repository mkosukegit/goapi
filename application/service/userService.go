package service

import (
	"context"
	"sample/application/repository"
	"sample/domain/entity"
)

type UserService interface {
	SelectAllUser(ctx context.Context) (user entity.User, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (userService *userService) SelectAllUser(ctx context.Context) (user entity.User, err error) {
	// userService.userRepository.SelectAllUser(ctx)
	user = entity.User{Id: 1, Name: "aa"}
	return
}
