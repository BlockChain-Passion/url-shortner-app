package controller

import (
	"html/template"
	"log"
	"net/http"
)

func ShowIndex(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/view/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

}
