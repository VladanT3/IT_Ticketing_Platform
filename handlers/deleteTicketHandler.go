package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
	"github.com/go-chi/chi/v5"
)

func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) error {
	ticketID := chi.URLParam(r, "ticketID")

	models.DeleteTicket(ticketID)

	return Render(w, r, tickets.DeletedTicket())
}
