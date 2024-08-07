package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
	"github.com/google/uuid"
)

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) error {
	typeCookie, err := r.Cookie("ticketType")
	if err != nil {
		log.Fatal("No cookie with name 'ticketType': ", err)
	}
	ticketType := typeCookie.Value

	categoryCookie, err := r.Cookie("category")
	if err != nil {
		log.Fatal("No cookie with name 'category': ", err)
	}
	category := categoryCookie.Value

	subcategoryCookie, err := r.Cookie("subcategory")
	if err != nil {
		log.Fatal("No cookie with name 'subcategory': ", err)
	}
	subcategory := subcategoryCookie.Value

	titleCookie, err := r.Cookie("title")
	if err != nil {
		log.Fatal("No cookie with name 'title': ", err)
	}
	title := titleCookie.Value

	descCookie, err := r.Cookie("desc")
	if err != nil {
		log.Fatal("No cookie with name 'desc': ", err)
	}
	desc := descCookie.Value

	customerContactCookie, err := r.Cookie("customerContact")
	if err != nil {
		log.Fatal("No cookie with name 'customerContact': ", err)
	}
	customerContact := customerContactCookie.Value

	saveTypeCookie, err := r.Cookie("saveType")
	if err != nil {
		log.Fatal("No cookie with name 'saveType': ", err)
	}
	saveType := saveTypeCookie.Value

	teamID := models.GetAnalystsTeam(LoggedInUser.Analyst_ID.String()).Team_ID

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

		return Render(w, r, tickets.TicketForm(models.Ticket{}, LoggedInUser, LoggedInUserType, "create", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
	}

	var db *sql.DB = database.DB_Connection
	query := `insert into ticket values(gen_random_uuid(), default, $1, 'Open', $2, $3, $4, $5, $6, default, default, null, $7, $8, $9, null) returning ticket_id;`
	newTicketID := ""
	err = db.QueryRow(query, ticketType, category, subcategory, title, desc, customerContact, teamID, LoggedInUser.Analyst_ID, LoggedInUser.Analyst_ID).Scan(&newTicketID)
	if err != nil {
		errMsg := "error inserting ticket: " + err.Error()
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
