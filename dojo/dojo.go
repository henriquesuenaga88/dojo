package dojo

import (
	"dojo/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type Dojo struct {
	Id    int
	Title string
	Done  bool
}

func (dojo *Dojo) isValid() bool {
	if dojo.Title != "" {
		return true
	}

	return false
}

type DojoPageData struct {
	PageTitle string
	Dojos     []Dojo
}

type DojoHandler struct {
	Db *mgo.Collection
}

func (h *DojoHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	var dojos []Dojo
	h.Db.Find(bson.M{}).All(&dojos)

	data := DojoPageData{
		PageTitle: "Dojo's",
		Dojos:     dojos,
	}
	render.Render(w, "index", data)
}

func (h *DojoHandler) NewDojo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		render.Render(w, "dojo/dojo", nil)
	} else {
		r.ParseForm()

		dojo := Dojo{Id: 0, Title: r.Form["title"][0], Done: false}
		if dojo.isValid() {
			h.Db.Insert(dojo)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			render.Render(w, "dojo/dojo", nil)
		}
	}
}

func (h *DojoHandler) ShowDetail(w http.ResponseWriter, r *http.Request) {
	dojo := Dojo{}
	var identificador, _ = strconv.ParseInt(r.URL.Query()["dojo"][0], 0, 64)
	h.Db.Find(bson.M{"id" : identificador}).One(&dojo)

	// TODO
}