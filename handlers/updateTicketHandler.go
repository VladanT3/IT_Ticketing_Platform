package handlers

import (
	"database/sql"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/ticket"
	"github.com/google/uuid"
)

func UpdateTicketHandler(w http.ResponseWriter, r *http.Request) error {
	ticketType := r.FormValue("type")
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	title := r.FormValue("title")
	desc := r.FormValue("desc")
	customerContact := r.FormValue("customerContact")
	saveType := r.FormValue("save")
	ticketID := r.FormValue("ticketID")

	if category == "none" {
		newTicket := models.Ticket{
			Type: ticketType,
			Category: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
			Subcategory: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
			Title:            title,
			Description:      desc,
			Customer_Contact: customerContact,
		}

		return Render(w, r, ticket.Ticket(models.GetTicket(ticketID), LoggedInUserType, "update", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
	}

	var db *sql.DB = database.DB_Connection
	query := `
		update ticket set
		type = $1,
		category_id = $2,
		subcategory_id = $3,
		title = $4,
		description = $5,
		customer_contact = $6,
		updated_at = current_timestamp
		returning ticket_id;
	`
	newTicketID := ""
	err := db.QueryRow(query, ticketType, category, subcategory, title, desc, customerContact).Scan(&newTicketID)
	if err != nil {
		errMsg := "error updating ticket: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	if saveType == "Save" {
		http.Redirect(w, r, "/ticket/"+newTicketID, http.StatusSeeOther)
		return nil
	} else if saveType == "Save and Exit" {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	return nil
}
