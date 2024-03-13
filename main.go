package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=server password=server dbname=userdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Exec("INSERT INTO usertable (id, login, password) VALUES (2, 'user1', 'pass1')")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.RowsAffected())
}
