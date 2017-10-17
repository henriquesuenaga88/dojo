package main

import (
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
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

func main() {
	http.HandleFunc("/", loadAll)
	http.HandleFunc("/dojo", dojo)

	http.ListenAndServe(":3000", nil)
}

func loadAll(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/layout.html"))

	session := getSession()
	defer session.Close()

	c := session.DB("dojo").C("dojo")
	var dojos [] Dojo
	c.Find(bson.M{}).All(&dojos)

	data := DojoPageData{
		PageTitle: "Dojo's",
		Dojos: dojos,
	}
	tmpl.Execute(w, data)

}

func dojo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("pages/dojo.html"))
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("title:", r.Form["title"])

		session := getSession()
		defer session.Close()

		c := session.DB("dojo").C("dojo")
		c.Insert(Dojo{Id: 0, Title: r.Form["title"][0], Done: false})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal(err)
	}
	return session
}
