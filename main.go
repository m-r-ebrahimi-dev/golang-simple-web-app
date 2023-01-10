package main

import (
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
