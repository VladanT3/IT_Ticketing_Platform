package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/ticket"
	"github.com/go-chi/chi/v5"
)

func ShowTicketHandler(w http.ResponseWriter, r *http.Request) error {
	ticketID := chi.URLParam(r, "ticketID")

	ticketToShow := models.GetTicket(ticketID)

	return Render(w, r, ticket.Ticket(ticketToShow, LoggedInUserType, "update", "", "", models.Ticket{}))
}
