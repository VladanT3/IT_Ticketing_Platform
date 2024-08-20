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
	router.Get("/get_subcategories", handlers.Make(handlers.SelectSubcategories))
	router.Post("/ticket", handlers.Make(handlers.TicketHandler))
	router.Get("/ticket/new", handlers.Make(handlers.NewTicketHandler))
	router.Get("/ticket/create", handlers.Make(handlers.CreateTicketHandler))
	router.Get("/ticket/update", handlers.Make(handlers.UpdateTicketHandler))
	router.Delete("/ticket/delete/{ticketID}", handlers.Make(handlers.DeleteTicketHandler))
	router.Get("/ticket/{ticketID}", handlers.Make(handlers.ShowTicketHandler))
	router.Get("/tickets/search", handlers.Make(handlers.ShowAllTicketSearchHandler))
	router.Get("/tickets/team", handlers.Make(handlers.ShowTeamTicketsHandler))
	router.Get("/tickets/unassigned", handlers.Make(handlers.ShowUnassignedTicketsHandler))
	router.Post("/search_tickets", handlers.Make(handlers.TicketSearchHandler))
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
