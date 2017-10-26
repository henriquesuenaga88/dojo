package main

import (
	"html/template"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
	"github.com/creamdog/gonfig"
	"os"
)

type Config struct {
	Database string
}

type DojoHandler struct {
	db *mgo.Collection
}

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
	config := loadConfig()

	session := config.getSession()
	defer session.Close()

	h := DojoHandler{db: session.DB("dojo").C("dojo")}

	http.HandleFunc("/", h.loadAll)
	http.HandleFunc("/dojo", h.dojo)

	http.ListenAndServe(":3000", nil)
}

func loadConfig() Config {
	f, _ := os.Open("properties.yml")
	defer f.Close()
	config, _ := gonfig.FromYml(f)
	database, _ := config.GetString("database", "")
	return Config{Database: database}
}

func (h *DojoHandler) loadAll(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/layout.html"))

	var dojos [] Dojo
	h.db.Find(bson.M{}).All(&dojos)

	data := DojoPageData{
		PageTitle: "Dojo's",
		Dojos: dojos,
	}
	tmpl.Execute(w, data)

}

func (h *DojoHandler) dojo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("pages/dojo.html"))
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("title:", r.Form["title"])

		h.db.Insert(Dojo{Id: 0, Title: r.Form["title"][0], Done: false})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (config *Config) getSession() *mgo.Session {
	session, err := mgo.Dial(config.Database)

	if err != nil {
		log.Fatal(err)
	}
	return session
}
