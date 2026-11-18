package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// connect to the postgres db
func Connect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// test the conn
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully")
	return db, nil
}
