package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	connStr := "user=postgres dbname=business sslmode=disable password=TaliusPostgres03"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening PostgreSQL: %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting PostgreSQL: %q", err)
	}

	fmt.Println("Connecting to PostgreSQL")
	return db, nil
}
