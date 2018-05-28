package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vaskoz/gophercises-phone"
)

func main() {
	dataSourceName := os.Getenv("MARIADBSQL_DATASOURCE")
	db, openerr := sql.Open("mysql", dataSourceName)
	if openerr != nil {
		log.Fatal(openerr)
	}
	for backoff := 1; backoff < 20; backoff *= 2 {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Duration(backoff) * time.Second)
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
