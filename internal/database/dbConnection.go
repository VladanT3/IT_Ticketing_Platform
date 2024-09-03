package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_Connection *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables: ", err)
	}

	db_conn := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", db_conn)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to ping database: ", err)
	}

	DB_Connection = db
}
