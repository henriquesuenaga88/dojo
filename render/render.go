package render

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
)
func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl_list := []string{"templates/layout.html",
		fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}