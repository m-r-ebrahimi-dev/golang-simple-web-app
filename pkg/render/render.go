package render

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using "html/template"
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("error parsing temlate: ", err)
		return
	}
}
