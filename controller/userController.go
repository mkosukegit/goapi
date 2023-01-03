package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"sample/application/service"
)

// 逆にDIされやすいようにインターフェースを作成
type UserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

// serviceをDI(コンストラクタインジェクション)
func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}

func (userController *userController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := userController.userService.SelectAllUser(context.Background())
	if err != nil {
		w.WriteHeader(500)
		return
	}

	output, _ := json.MarshalIndent(users, "", "\t\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
