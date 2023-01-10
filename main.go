package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const portNumber = ":8080"

// HomePage is about index page
func HomePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Home Page")
	if err != nil {
		log.Fatal(err)
		return
	}
}

// About is intro page
func About(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	sum, _ := addValues(rand.Intn(10), rand.Intn(10))
	_, err := fmt.Fprintf(w, fmt.Sprintf("about %d", sum))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func addValues(x, y int) (int, error) {
	sum := x + y
	return sum, nil
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
