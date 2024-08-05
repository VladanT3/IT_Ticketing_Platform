package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
)

func TicketSearchHandler(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("search")
	customer := r.FormValue("customer")
	ticketType := r.FormValue("type")
	status := r.FormValue("status")
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	searchType := r.FormValue("searchType")

	tickets := models.FilterTickets(search, customer, ticketType, status, category, subcategory, searchType, LoggedInUser.Team_ID.UUID.String())

	return Render(w, r, layouts.Tickets(tickets))
}
