package middleware

import (
	"net/http"
	"sample/controller"
)

type Middleware interface {
	Connect()
}

func Connect() {
	http.HandleFunc("/api/user/", controller.Routing)
}
