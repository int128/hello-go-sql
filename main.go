package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func run() error {
	db, err := sql.Open("postgres", "postgres://app:app123@localhost/postgres")
	if err != nil {
		return fmt.Errorf("could not open connection: %w", err)
	}
	defer db.Close()

	result, err := db.Exec(`
SELECT 1;
SELECT 2;
`)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rowsAffected: %w", err)
	}
	log.Printf("%d rows affected", rowsAffected)

	return nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	if err := run(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
