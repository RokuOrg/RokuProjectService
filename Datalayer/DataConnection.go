package Datalayer

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func Connect() {
	cfg := mysql.Config{
		User:   os.Getenv("ROKU_AUTH_USERNAME"),
		Passwd: os.Getenv("ROKU_AUTH_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("ROKU_AUTH_IP"),
		DBName: "RokuProjects",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
