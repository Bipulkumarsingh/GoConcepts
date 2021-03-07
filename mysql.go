package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mainsql() {
	fmt.Println("Connecting with mysql")

	// Open up our database connection.
	db, err := sql.Open("mysql", "")
}
