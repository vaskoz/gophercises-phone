package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

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
		fmt.Printf("%s: %s\n", number, Normalize(number))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
