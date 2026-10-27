package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Customer struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
}

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `SELECT * FROM customer ORDER BY name ASC`
	rows, err := db.Query(query, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	customers := []Customer{}

	for rows.Next() {
		var c Customer

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Address,
			&c.City,
			&c.Zip,
		)
		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	jsonOut, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("customers.json", jsonOut, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
