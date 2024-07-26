package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/VladanT3/IT_Ticketing_Platform/handlers"
	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	database.Connect()
	defer database.DB_Connection.Close()

	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	router.Get("/", handlers.Make(handlers.IndexHandler))
	router.Post("/login", handlers.Make(handlers.LoginHandler))
	router.Post("/logout", handlers.Make(handlers.LogoutHandler))
	router.Get("/profile", handlers.Make(handlers.ProfileHandler))
	router.Get("/ticket/new", handlers.Make(handlers.NewTicketHandler))

	port := os.Getenv("PORT")
	fmt.Println("Server started on: http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
