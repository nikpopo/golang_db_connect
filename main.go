package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                     "un",
		Passwd:                   "pw",
		Net:                      "tcp",
		Addr:                     "localhost:3306",
		DBName:                   "db",
		Collation: "utf8mb4_general_ci",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
}
