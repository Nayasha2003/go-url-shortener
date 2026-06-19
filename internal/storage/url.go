package storage

import (
	"database/sql"
	"log"
)

func InsertURL(short, original string) error {
	_, err := DB.Exec(
		"INSERT INTO urls (short_code, original_url) VALUES ($1, $2)",
		short, original,
	)
	if err != nil {
		log.Println("InsertURL error:", err)
	}
	return err
}

func GetOriginalURL(short string) (string, error) {
	var original string
	err := DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_code = $1", short,
	).Scan(&original)
	if err == sql.ErrNoRows {
		return "", err
	} else if err != nil {
		log.Println("GetOriginalURL error:", err)
		return "", err
	}
	return original, nil
}
