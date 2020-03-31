package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	postgresDSN := os.Getenv("POSTGRES_DSN")

	for {
		db, err := NewPostgres(postgresDSN)
		if err != nil {
			continue
		}
		db.Close()
		break
	}

	fmt.Println("Up")

	os.Exit(0)
}

// NewPostgres creates a new Database with postgres as driver.
func NewPostgres(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
