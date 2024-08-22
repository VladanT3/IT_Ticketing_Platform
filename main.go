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
	//TODO: fix the border styling somehow for ticket status
	//TODO: update and track the numbers of open/opened/closed tickets
	router.Get("/profile", handlers.Make(handlers.Profile))

	router.Post("/ticket", handlers.Make(handlers.TicketRedirection))
	router.Get("/ticket/new", handlers.Make(handlers.ShowNewTicketForm))
	router.Get("/ticket/create", handlers.Make(handlers.CreateTicket))
	router.Get("/ticket/update", handlers.Make(handlers.UpdateTicket))
	router.Delete("/ticket/delete/{ticketID}", handlers.Make(handlers.DeleteTicket))
	router.Get("/ticket/{ticketID}", handlers.Make(handlers.ShowTicket))
	router.Put("/ticket/close/{ticketID}", handlers.Make(handlers.CloseTicket))
	//TODO: implement ticket reopen

	router.Get("/tickets/search", handlers.Make(handlers.ShowAllTicketSearch))
	router.Get("/tickets/team", handlers.Make(handlers.ShowTeamTickets))
	router.Get("/tickets/unassigned", handlers.Make(handlers.ShowUnassignedTickets))
	router.Post("/tickets/filter", handlers.Make(handlers.FilterTickets))

	router.Get("/categories", handlers.Make(handlers.ShowCategoriesPage))
	router.Get("/categories/search", handlers.Make(handlers.SearchCategories))
	router.Get("/category/popup", handlers.Make(handlers.ShowCategoryPopup))
	router.Get("/category/error/name", handlers.Make(handlers.ShowCategoryAlreadyExistsError))
	router.Post("/category/create", handlers.Make(handlers.CreateCategory))
	router.Put("/category/update/{categoryID}", handlers.Make(handlers.UpdateCategory))
	router.Delete("/category/delete/{categoryID}", handlers.Make(handlers.DeleteCategory))

	router.Get("/subcategories/get/select", handlers.Make(handlers.SelectSubcategories))
	router.Get("/subcategories/get/modifiable", handlers.Make(handlers.ShowModifiableSubcategories))
	router.Get("/subcategories/search", handlers.Make(handlers.SearchSubcategories))
	router.Get("/subcategory/popup", handlers.Make(handlers.ShowSubcategoryPopup))
	router.Get("/subcategory/error/name", handlers.Make(handlers.ShowSubcategoryAlreadyExistsError))
	router.Post("/subcategory/create", handlers.Make(handlers.CreateSubcategory))
	router.Put("/subcategory/update/{subcategoryID}", handlers.Make(handlers.UpdateSubcategory))
	router.Delete("/subcategory/delete/{subcategoryID}", handlers.Make(handlers.DeleteSubcategory))

	port := os.Getenv("PORT")
	fmt.Println("Server started on: http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
