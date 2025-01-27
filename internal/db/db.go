package db

import (
	"database/sql"
	"errors"
	"log"
)

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
    	ID  SERIAL PRIMARY KEY,
    	short_url VARCHAR(255) NOT NULL,
    	origin_url VARCHAR(255) NOT NULL
			);`
	log.Println("We are here at CreateTable")
	_, err := db.Exec(query)
	log.Println("We are here at CreateTable....1", err)
	return err
}

func StoreURL(db *sql.DB, shortURL string, originURL string) error {
	query := `INSERT INTO urls (short_url, origin_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originURL)
	return err
}

func GetOriginURL(db *sql.DB, shortURL string) (string, error) {
	var originURL string
	log.Println("We are here at GetOriginURL", shortURL)
	query := `SELECT origin_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortURL).Scan(&originURL)
	if err != nil {
		return "", errors.New("not found")
	}
	return originURL, nil
}
