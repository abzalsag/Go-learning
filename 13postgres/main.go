package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	connStr := "host=localhost port=54320 user=db_user password=pwd123 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to DB")
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Server started")

	http.ListenAndServe(":8080", nil)
}
