package dojo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
)

type Dojo struct {
	Id    int
	Title string
	Done  bool
}

type DojoPageData struct {
	PageTitle string
	Dojos     []Dojo
}

type DojoHandler struct {
	Db *mgo.Collection
}

func (h *DojoHandler) FindAll(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	var dojos []Dojo
	h.Db.Find(bson.M{}).All(&dojos)

	data := DojoPageData{
		PageTitle: "Dojo's",
		Dojos:     dojos,
	}
	tmpl.Execute(w, data)
}

func (h *DojoHandler) NewDojo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("templates/dojo/dojo.html"))
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("title:", r.Form["title"])

		h.Db.Insert(Dojo{Id: 0, Title: r.Form["title"][0], Done: false})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

