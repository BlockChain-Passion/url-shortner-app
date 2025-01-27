package controller

import (
	"database/sql"
	"github.com/BlockChain-Passion/url-shortner-app/internal/db"
	"github.com/BlockChain-Passion/url-shortner-app/internal/url"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		originalURL := r.FormValue("url")
		if originalURL == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !strings.HasPrefix(originalURL, "https://") && !strings.HasPrefix(originalURL, "http://") {
			originalURL = "https://" + originalURL
		}

		shortURL := url.Shorten(originalURL)

		if err := db.StoreURL(lite, shortURL, originalURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//shorten  url
		data := map[string]string{
			"ShortURL": shortURL,
		}

		t, err := template.ParseFiles("internal/view/shorten.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func Redirect(lite *sql.DB) http.HandlerFunc {
	log.Println("Redirecting method")
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:]
		log.Println("In Redirecting method")
		log.Println(shortURL)

		if shortURL == "" {
			http.Error(w, "shorten url not found", http.StatusNotFound)
			return
		}
		origUrl, err := db.GetOriginURL(lite, shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, origUrl, http.StatusPermanentRedirect)
	}

}
