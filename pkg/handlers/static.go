package handlers

import "net/http"

var StaticHandler = http.StripPrefix(
	"/static/",
	http.FileServer(http.Dir("./static")),
)
