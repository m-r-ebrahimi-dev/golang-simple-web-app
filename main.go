package main

import (
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

// HomePage is about index page
func HomePage(w http.ResponseWriter, r *http.Request) {
	/*_, err := fmt.Fprintf(w, "Home Page")
	if err != nil {
		log.Fatal(err)
		return
	}*/
	renderTemplate(w, "home.page.tmpl")

}

// About is intro page
func About(w http.ResponseWriter, r *http.Request) {
	/*rand.Seed(time.Now().UnixNano())
	sum, _ := addValues(rand.Intn(10), rand.Intn(10))
	_, err := fmt.Fprintf(w, fmt.Sprintf("about %d", sum))
	if err != nil {
		log.Fatal(err)
		return
	}*/
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("error parsing temlate: ", err)
		return
	}

}

/*func Divide(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	divide, error := divideValues(rand.Float32(), rand.Float32())
	if error != nil {
		fmt.Fprintf(w, fmt.Sprintf("can't divide by 0"))
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("divide: %.2f", divide))

}*/

/*func addValues(x, y int) (int, error) {
	sum := x + y
	return sum, nil
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("divider must not be 0")
	}
	return x / y, nil
}*/

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", About)
	//http.HandleFunc("/divide", Divide)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
