package main

import (
	"github.com/m-r-ebrahimi-dev/golang-simple-web-app/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.About)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
