package server

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "layout", "header", "index", "footer")
}
