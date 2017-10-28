package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/creamdog/gonfig"
	"os"
	"dojo/dojo"
)

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
