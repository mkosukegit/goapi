package controller

import (
	"net/http"
	"sample/application/repository"
	"sample/application/service"
)

type UserRouter interface {
	Routing(w http.ResponseWriter, r *http.Request)
}

func Routing(w http.ResponseWriter, r *http.Request) {

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := NewUserController(userService)

	switch r.Method {
	case "GET":
		userController.GetUsers(w, r)
	default:
		w.WriteHeader(405)
	}

}
