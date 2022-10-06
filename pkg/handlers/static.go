package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("./static/html/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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
