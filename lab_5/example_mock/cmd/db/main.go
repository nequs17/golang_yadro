package main

import (
	"database/sql"
	dbPack "example_mock/internal/db"
	"fmt"
	"log"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()

	for _, name := range names {
		fmt.Println(name)
	}
}
