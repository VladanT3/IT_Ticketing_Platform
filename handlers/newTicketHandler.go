package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/ticket"
)

func NewTicketHandler(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, ticket.Ticket(models.Ticket{}, LoggedInUserType, "create"))
}
