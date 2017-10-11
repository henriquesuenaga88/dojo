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
	tmpl := template.Must(template.ParseFiles("pages/layout.html"))

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("dojo").C("dojo")


	count, err2 := c.Find(bson.M{}).Count()
	if err2 != nil {
		panic(err)
	}

	// FOR TEST
	if count == 0 {
		dojoteste1 := Dojo{Id: 1, Title: "Teste 1", Done: false}
		c.Insert(dojoteste1)

		dojoteste2 := Dojo{Id: 2, Title: "Teste 2", Done: false}
		c.Insert(dojoteste2)

		dojoteste3 := Dojo{Id: 3, Title: "Teste 3", Done: false}
		c.Insert(dojoteste3)

		dojoteste4 := Dojo{Id: 4, Title: "Teste 4", Done: false}
		c.Insert(dojoteste4)
	}

	var dojos [] Dojo
	c.Find(bson.M{}).All(&dojos)

	fmt.Println(dojos[0].Title)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := DojoPageData{
			PageTitle: "Dojo's",
			Dojos: dojos,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":3000", nil)
}
