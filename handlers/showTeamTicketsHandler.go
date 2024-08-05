package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
)

func ShowTeamTicketsHandler(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "Team Tickets"))
}