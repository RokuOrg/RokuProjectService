package main

import (
	"RokuProject-Back-End/Program"
	"github.com/go-sql-driver/mysql"
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

	a := Program.BuildApp(cfg)

	a.Run(":7001")
}
