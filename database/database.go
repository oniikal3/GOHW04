package database

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func main() {}

func Open() *sql.DB {
	newDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	db = newDB
	return newDB
}

func Close() {
	db.Close()
}
