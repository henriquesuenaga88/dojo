package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/creamdog/gonfig"
	"os"
	"dojo/dojo"
	"io"
	"time"
)


const STATIC_URL string = "/static/"
const STATIC_ROOT string = "static/"

type Config struct {
	Database string
}

func main() {
	config := loadConfig()

	session := config.getSession()
	defer session.Close()

	h := dojo.DojoHandler{Db: session.DB("dojo").C("dojo")}

	http.HandleFunc("/", h.FindAll)
	http.HandleFunc("/dojo", h.NewDojo)
	http.HandleFunc(STATIC_URL, StaticHandler)

	http.ListenAndServe(":3000", nil)
}

func loadConfig() Config {
	f, _ := os.Open("properties.yml")
	defer f.Close()
	config, _ := gonfig.FromYml(f)
	database, _ := config.GetString("database", "")
	return Config{Database: database}
}

func (config *Config) getSession() *mgo.Session {
	session, err := mgo.Dial(config.Database)

	if err != nil {
		log.Fatal(err)
	}
	return session
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}