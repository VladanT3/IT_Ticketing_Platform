package handlers

import (
	"database/sql"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
	"github.com/google/uuid"
)

func UpdateTicketHandler(w http.ResponseWriter, r *http.Request) error {
	typeCookie, err := r.Cookie("ticketType")
	if err != nil {
		errMsg := "No cookie with name 'ticketType': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	ticketType := typeCookie.Value

	categoryCookie, err := r.Cookie("category")
	if err != nil {
		errMsg := "No cookie with name 'category': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	category := categoryCookie.Value

	subcategoryCookie, err := r.Cookie("subcategory")
	if err != nil {
		errMsg := "No cookie with name 'subcategory': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	subcategory := subcategoryCookie.Value

	titleCookie, err := r.Cookie("title")
	if err != nil {
		errMsg := "No cookie with name 'title': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	title := titleCookie.Value

	descCookie, err := r.Cookie("desc")
	if err != nil {
		errMsg := "No cookie with name 'desc': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	desc := descCookie.Value

	customerContactCookie, err := r.Cookie("customerContact")
	if err != nil {
		errMsg := "No cookie with name 'customerContact': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	customerContact := customerContactCookie.Value

	saveTypeCookie, err := r.Cookie("saveType")
	if err != nil {
		errMsg := "No cookie with name 'saveType': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	saveType := saveTypeCookie.Value

	ticketIDCookie, err := r.Cookie("ticketID")
	if err != nil {
		errMsg := "No cookie with name 'ticketID': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	ticketID := ticketIDCookie.Value

	if category == "none" || category == "" {
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

		return Render(w, r, tickets.Ticket(models.GetTicket(ticketID), LoggedInUser, LoggedInUserType, "update", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
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
		where ticket_id = $7
		returning ticket_id;
	`
	newTicketID := ""
	err = db.QueryRow(query, ticketType, category, subcategory, title, desc, customerContact, ticketID).Scan(&newTicketID)
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
