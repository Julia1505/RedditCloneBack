package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("./static/html/index.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Println("err: %w", err)
	}
	log.Println("OK")
}

var StaticHandler = http.StripPrefix(
	"/static/",
	http.FileServer(http.Dir("./static")),
)
