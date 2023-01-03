package controller

import (
	"fmt"
	"net/http"
)

func bookController(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "bookController")
}
