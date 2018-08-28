package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vaskoz/gophercises-phone"
)

var (
	stdout         io.Writer = os.Stdout
	exit           func(int) = os.Exit
	dataSourceName string    = os.Getenv("MARIADBSQL_DATASOURCE")
)

func reset() {
	stdout = os.Stdout
	exit = os.Exit
	dataSourceName = os.Getenv("MARIADBSQL_DATASOURCE")
}

func main() {
	logger := log.New(stdout, "mariadbsql: ", log.Lshortfile)
	db, _ := sql.Open("mysql", dataSourceName)
	if err := db.Ping(); err != nil {
		logger.Println(err)
		exit(1)
		return
	}
	rows, _ := db.Query("SELECT number FROM phone")
	for rows.Next() {
		var number string
		rows.Scan(&number)
		logger.Printf("%s: %s\n", number, phone.Normalize(number))
	}
	exit(0)
}
