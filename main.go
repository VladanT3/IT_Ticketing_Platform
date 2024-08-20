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
	router.Get("/", handlers.Make(handlers.Index))
	router.Post("/login", handlers.Make(handlers.Login))
	router.Post("/logout", handlers.Make(handlers.Logout))
	router.Get("/profile", handlers.Make(handlers.Profile))
	router.Get("/get_subcategories", handlers.Make(handlers.SelectSubcategories))
	router.Post("/ticket", handlers.Make(handlers.TicketRedirection))
	router.Get("/ticket/new", handlers.Make(handlers.ShowNewTicketForm))
	router.Get("/ticket/create", handlers.Make(handlers.CreateTicket))
	router.Get("/ticket/update", handlers.Make(handlers.UpdateTicket))
	router.Delete("/ticket/delete/{ticketID}", handlers.Make(handlers.DeleteTicket))
	router.Get("/ticket/{ticketID}", handlers.Make(handlers.ShowTicket))
	router.Get("/tickets/search", handlers.Make(handlers.ShowAllTicketSearch))
	router.Get("/tickets/team", handlers.Make(handlers.ShowTeamTickets))
	router.Get("/tickets/unassigned", handlers.Make(handlers.ShowUnassignedTickets))
	router.Post("/search_tickets", handlers.Make(handlers.TicketSearch))
	router.Get("/categories", handlers.Make(handlers.ShowCategoriesPage))
	router.Get("/show_subcategories", handlers.Make(handlers.EditingSubcategories))
	router.Get("/search_categories", handlers.Make(handlers.SearchCategories))
	router.Get("/search_subcategories", handlers.Make(handlers.SearchSubcategories))
	router.Get("/category/create/popup", handlers.Make(handlers.ShowCategoryPopup))
	router.Get("/category/create", handlers.Make(handlers.))
	//router.Get("/subcategory/create", handlers.Make())
	//router.Get("/category/edit/{categoryID}", handlers.Make())
	//router.Get("/subcategory/edit/{categoryID}", handlers.Make())

	port := os.Getenv("PORT")
	fmt.Println("Server started on: http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
