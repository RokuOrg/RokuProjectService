package Datalayer

import (
	"database/sql"
)

type DatabaseLayer struct {
	DB *sql.DB
}

func CreateDatabaseLayer(db *sql.DB) *DatabaseLayer {
	dl := DatabaseLayer{DB: db}
	return &dl
}
