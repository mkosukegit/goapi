package main

import (
	"log"
	"net/http"
	"sample/middleware"
)

func main() {

	middleware.Connect()

	print("サーバ-を起動します")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
