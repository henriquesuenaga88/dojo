package main

import (
	"html/template"
	"net/http"
)

type Dojo struct {
	Title string
	Done  bool
}

type DojoPageData struct {
	PageTitle string
	Dojos     []Dojo
}

func main() {
	tmpl := template.Must(template.ParseFiles("pages/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := DojoPageData{
			PageTitle: "Dojo's",
			Dojos: []Dojo{
				{Title: "Dojo 1", Done: false},
				{Title: "Dojo 2", Done: true},
				{Title: "Dojo 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":3000", nil)
}
