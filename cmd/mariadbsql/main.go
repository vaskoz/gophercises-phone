package main

import (
	"database/sql"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vaskoz/gophercises-phone"
)

var (
	stdout io.Writer = os.Stdout
	exit   func(int) = os.Exit
)

func main() {
	dataSourceName := os.Getenv("MARIADBSQL_DATASOURCE")
	logger := log.New(stdout, "mariadbsql: ", log.Lshortfile)
	db, openerr := sql.Open("mysql", dataSourceName)
	if openerr != nil {
		logger.Println(openerr)
		exit(1)
		return
	}
	for backoff := 1; backoff < 20; backoff *= 2 {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	rows, err := db.Query("SELECT number FROM phone")
	if err != nil {
		logger.Println(err)
		exit(1)
		return
	}
	for rows.Next() {
		var number string
		if err := rows.Scan(&number); err != nil {
			logger.Println(err)
			exit(1)
			return
		}
		logger.Printf("%s: %s\n", number, phone.Normalize(number))
	}
	if err := rows.Err(); err != nil {
		logger.Println(err)
		exit(1)
		return
	}
}
