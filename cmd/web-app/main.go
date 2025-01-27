package main

import (
	"database/sql"
	"github.com/BlockChain-Passion/url-shortner-app/internal/controller"
	"github.com/BlockChain-Passion/url-shortner-app/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	slide, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	//defer func() {
	//	err := slide.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	if err = db.CreateTable(slide); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controller.ShowIndex(w, r)
		} else if r.URL.Path != "/favicon.ico" {
			log.Printf("We are here ")
			controller.Redirect(slide)
		}
	})
	http.HandleFunc("/shorten", controller.Shorten(slide))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
