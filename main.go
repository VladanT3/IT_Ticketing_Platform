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

	router.Post("/ticket", handlers.Make(handlers.TicketRedirection))
	router.Get("/ticket/new", handlers.Make(handlers.ShowNewTicketForm))
	router.Get("/ticket/create", handlers.Make(handlers.CreateTicket))
	router.Get("/ticket/update", handlers.Make(handlers.UpdateTicket))
	router.Delete("/ticket/delete/{ticket_id}", handlers.Make(handlers.DeleteTicket))
	router.Get("/ticket/{ticket_id}", handlers.Make(handlers.ShowTicket))
	router.Put("/ticket/close/{ticket_id}", handlers.Make(handlers.CloseTicket))
	router.Get("/ticket/{ticket_id}/reopen/form", handlers.Make(handlers.ShowTicketReopenForm))
	router.Post("/ticket/{ticket_id}/reopen", handlers.Make(handlers.ReopenTicket))
	router.Get("/ticket/{ticket_id}/reopen/history", handlers.Make(handlers.ShowTicketReopenHistory))
	router.Get("/ticket/{ticket_id}/assign/form", handlers.Make(handlers.ShowTicketAssignmentForm))
	router.Post("/ticket/{ticket_id}/assign", handlers.Make(handlers.AssignTicket))
	router.Get("/ticket/{ticket_id}/assign/history", handlers.Make(handlers.ShowTicketAssignmentHistory))
	router.Post("/ticket/{ticket_id}/assign/self", handlers.Make(handlers.AssignTicketToMe))

	router.Get("/tickets/search", handlers.Make(handlers.ShowAllTicketSearch))
	router.Get("/tickets/team", handlers.Make(handlers.ShowTeamTickets))
	router.Get("/tickets/unassigned", handlers.Make(handlers.ShowUnassignedTickets))
	router.Post("/tickets/filter", handlers.Make(handlers.FilterTickets))

	router.Get("/categories", handlers.Make(handlers.ShowCategoriesPage))
	router.Get("/categories/search", handlers.Make(handlers.SearchCategories))
	router.Get("/category/popup", handlers.Make(handlers.ShowCategoryPopup))
	router.Get("/category/error/name", handlers.Make(handlers.ShowCategoryAlreadyExistsError))
	router.Post("/category/create", handlers.Make(handlers.CreateCategory))
	router.Put("/category/update/{category_id}", handlers.Make(handlers.UpdateCategory))
	router.Delete("/category/delete/{category_id}", handlers.Make(handlers.DeleteCategory))

	router.Get("/subcategories/get/select", handlers.Make(handlers.SelectSubcategories))
	router.Get("/subcategories/get/modifiable", handlers.Make(handlers.ShowModifiableSubcategories))
	router.Get("/subcategories/search", handlers.Make(handlers.SearchSubcategories))
	router.Get("/subcategory/popup", handlers.Make(handlers.ShowSubcategoryPopup))
	router.Get("/subcategory/error/name", handlers.Make(handlers.ShowSubcategoryAlreadyExistsError))
	router.Post("/subcategory/create", handlers.Make(handlers.CreateSubcategory))
	router.Put("/subcategory/update/{subcategory_id}", handlers.Make(handlers.UpdateSubcategory))
	router.Delete("/subcategory/delete/{subcategory_id}", handlers.Make(handlers.DeleteSubcategory))

	router.Get("/team/analysts", handlers.Make(handlers.GetTeamsAnalysts))
	router.Get("/analyst/team", handlers.Make(handlers.GetAnalystsTeam))

	router.Get("/users/view", handlers.Make(handlers.ShowUserView))
	router.Get("/users/team/view", handlers.Make(handlers.ShowTeamView))
	router.Post("/users/filter", handlers.Make(handlers.FilterUsers))

	router.Get("/user/{analyst_id}", handlers.Make(handlers.ShowUserForm))
	router.Get("/user/new", handlers.Make(handlers.ShowNewUserForm))
	router.Post("/user", handlers.Make(handlers.UserRedirect))
	router.Get("/user/create", handlers.Make(handlers.CreateUser))
	router.Get("/user/update/{analyst_id}", handlers.Make(handlers.UpdateUser))
	router.Delete("/user/delete/{analyst_id}", handlers.Make(handlers.DeleteUser))
	router.Get("/user/request/update/{analyst_id}", handlers.Make(handlers.RequestUserInfoChange))
	router.Get("/user/password/change/form", handlers.Make(handlers.ShowChangePasswordForm))
	router.Post("/user/password/change", handlers.Make(handlers.ChangePassword))
	router.Get("/user/password/change/success", handlers.Make(handlers.ShowPasswordChangeSuccess))

	router.Get("/error", handlers.Make(handlers.ShowError))

	port := os.Getenv("PORT")
	fmt.Println("Server started on: http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
