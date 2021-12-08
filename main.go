package main

import (
	"RokuProject-Back-End/Program"
	"github.com/go-sql-driver/mysql"
	"github.com/segmentio/ksuid"
	"os"
)

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("ROKU_AUTH_USERNAME"),
		Passwd: os.Getenv("ROKU_AUTH_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("ROKU_AUTH_IP"),
		DBName: "RokuProjects",
	}

	db := Program.CreateDbConnection(cfg)
	a := Program.BuildApp(db, ksuid.New)

	a.Run(":7001")
}
