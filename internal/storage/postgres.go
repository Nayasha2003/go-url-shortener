package storage

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Ctx = context.Background()
var DB *sql.DB

func InitPostgres() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:Omanand%402000@localhost:5432/urlshortener?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Postgres open error:", err)
	}

	if err = db.PingContext(Ctx); err != nil {
		log.Fatal("Postgres ping error:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")

	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
		short_code   VARCHAR(10) PRIMARY KEY,
		original_url TEXT NOT NULL,
		created_at   TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	if _, err := DB.ExecContext(Ctx, createTable); err != nil {
		log.Fatal("Failed to create urls table:", err)
	}

	log.Println("urls table ready")
}
