package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
	"github.com/go-chi/chi/v5"
)

func ShowTicketHandler(w http.ResponseWriter, r *http.Request) error {
	ticketID := chi.URLParam(r, "ticketID")

	if models.TicketExists(ticketID) {
		ticketToShow := models.GetTicket(ticketID)
		return Render(w, r, tickets.TicketForm(ticketToShow, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
	} else {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "No ticket found: ticket doesn't exist or it was removed!"))
	}
}
