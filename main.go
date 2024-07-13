package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/VladanT3/IT_Ticketing_Platform/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables: ", err)
	}

	dbConn := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to ping database: ", err)
	}

	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	router.Get("/", handlers.Make(handlers.LoginHandler))

	port := os.Getenv("PORT")
	fmt.Println("Server started on: http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
