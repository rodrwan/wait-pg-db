package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")

	for {
		db, err := NewPostgres(databaseURL)
		if err != nil {
			fmt.Println("waiting...")
			time.Sleep(1 * time.Second)
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
