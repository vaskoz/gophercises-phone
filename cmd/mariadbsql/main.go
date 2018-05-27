package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vaskoz/gophercises-phone"
)

func main() {
	dataSourceName := os.Getenv("MARIADBSQL_DATASOURCE")
	db, err := sql.Open("mysql", dataSourceName)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT number FROM phone")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var number string
		if err := rows.Scan(&number); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", number, phone.Normalize(number))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
