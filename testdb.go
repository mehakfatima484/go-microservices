package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=root123 host=localhost port=5433 dbname=go_tasks sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	fmt.Println("Successfully connected to the database.")
}
